package saver

import (
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
	s.dataContainer = append(s.dataContainer, entity)
	s.mutex.Unlock()
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

	copyData := append(s.dataContainer[:0:0], s.dataContainer...)

	if len(copyData) != 0 {
		s.flusher.Flush(copyData)
		s.dataContainer = nil
	}
}
