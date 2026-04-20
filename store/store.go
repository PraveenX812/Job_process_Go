package store

import (
	"go-jobs/models"
	"sync"
)

//thread safe in memory storage
type Store struct {
	mu   sync.RWMutex
	jobs map[string]*models.Job
}

func New() *Store {
	return &Store{
		jobs: make(map[string]*models.Job),
	}
}

//adds a new job to the store
func (s *Store) Save(job *models.Job) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs[job.ID] = job
}

//retrieve job from the store using id
func (s *Store) Get(id string) (*models.Job, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	job, ok := s.jobs[id]
	return job, ok
}

//updates an existing job in the store
func (s *Store) Update(job *models.Job) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs[job.ID] = job
}
