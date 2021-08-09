package saver

import (
	"github.com/ozoncp/ocp-timeline-api/internal/flusher"
	"github.com/ozoncp/ocp-timeline-api/internal/models"
	"time"
)

const duration = time.Second * 3

type Saver interface {
	Save(entity models.Timeline)
	Close()
}

func NewSaver(capacity int, flusher flusher.Flusher) Saver {
	saver := &saver{
		capacity:   capacity,
		flusher:    flusher,
		stopChanel: make(chan int),
		dataChanel: make(chan models.Timeline, capacity),
	}

	saver.init()

	return saver
}

type saver struct {
	capacity   int
	flusher    flusher.Flusher
	stopChanel chan int
	dataChanel chan models.Timeline
}

func (s *saver) Save(entity models.Timeline) {
	go func() {
		s.dataChanel <- entity
	}()
}

func (s *saver) Close() {
	s.stopChanel <- 1
}

func (s *saver) init() {
	var ticker = time.NewTicker(duration)
	defer ticker.Stop()
	defer close(s.dataChanel)
	defer close(s.stopChanel)

	go func() {
		for {
			select {
			case <-ticker.C:
				s.inputInFlusher()
			case <-s.stopChanel:
				for {
					if len(s.dataChanel) != 0 {
						s.inputInFlusher()
					} else {
						return
					}
				}
			}
		}
	}()
}

func (s *saver) readFromChanelInSlice() []models.Timeline {
	result := make([]models.Timeline, 0, s.capacity)

	for i := 0; i < len(s.dataChanel); i++ {
		result = append(result, <-s.dataChanel)
	}

	return result
}

func (s *saver) inputInFlusher() {
	data := s.readFromChanelInSlice()
	if len(data) != 0 {
		s.flusher.Flush(data)
	}
}
