package repository

import (
	"sync"
)

// UserService handles user language caching and DB
type UserService struct {
	repo  *UserRepository
	cache map[int64]string
	mu    sync.RWMutex
}

// NewUserService creates a new UserService
func NewUserService(repo *UserRepository) *UserService {
	return &UserService{
		repo:  repo,
		cache: make(map[int64]string),
	}
}

// GetUserLanguage returns the cached language, or DB fallback
func (s *UserService) GetUserLanguage(userID int64) string {
	s.mu.RLock()
	lang, ok := s.cache[userID]
	s.mu.RUnlock()
	if ok {
		return lang
	}

	// Fallback to DB
	lang, err := s.repo.GetUserLanguage(userID)
	if err != nil {
		lang = "en" // default
	}

	// Store in cache
	s.mu.Lock()
	s.cache[userID] = lang
	s.mu.Unlock()

	return lang
}

// SetUserLanguage updates cache and DB
func (s *UserService) SetUserLanguage(userID int64, lang string) error {
	// Update cache
	s.mu.Lock()
	s.cache[userID] = lang
	s.mu.Unlock()

	// Update DB
	return s.repo.CreateOrUpdateLanguage(userID, lang)
}
