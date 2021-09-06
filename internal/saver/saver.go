package saver

import (
	"context"
	"github.com/ozoncp/ocp-timeline-api/internal/flusher"
	"github.com/ozoncp/ocp-timeline-api/internal/models"
	"sync"
	"time"
)

const duration = time.Second * 3

type Saver interface {
	Save(entity models.Timeline)
	Close()
}

func NewSaver(capacity int, flusher flusher.Flusher) Saver {
	saver := &saver{
		capacity:      capacity,
		flusher:       flusher,
		stopChanel:    make(chan int),
		dataContainer: make([]models.Timeline, 0, capacity),
		mutex:         sync.Mutex{},
	}

	saver.init()

	return saver
}

type saver struct {
	capacity      int
	flusher       flusher.Flusher
	stopChanel    chan int
	dataContainer []models.Timeline
	mutex         sync.Mutex
}

func (s *saver) Save(entity models.Timeline) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.dataContainer = append(s.dataContainer, entity)
}

func (s *saver) Close() {
	s.stopChanel <- 1
}

func (s *saver) init() {
	var ticker = time.NewTicker(duration)

	go func() {
		defer ticker.Stop()
		defer close(s.stopChanel)

		for {
			select {
			case <-ticker.C:
				s.inputInFlusher()
			case <-s.stopChanel:
				s.inputInFlusher()
				return
			}
		}
	}()
}

func (s *saver) inputInFlusher() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.dataContainer) != 0 {
		s.flusher.Flush(context.TODO(), s.dataContainer)
		s.dataContainer = nil
	}
}
