package math

import (
	"sync"

	domainMath "github.com/ifabos/g4o/internal/domain/math"
)

// InMemoryMathRepository implements MathRepository interface using in-memory storage
type InMemoryMathRepository struct {
	operations []*domainMath.MathOperation
	results    []*domainMath.Result
	mutex      sync.RWMutex
}

// NewInMemoryMathRepository creates a new InMemoryMathRepository instance
func NewInMemoryMathRepository() *InMemoryMathRepository {
	return &InMemoryMathRepository{
		operations: make([]*domainMath.MathOperation, 0),
		results:    make([]*domainMath.Result, 0),
	}
}

// SaveOperation saves an operation and its result to memory
func (r *InMemoryMathRepository) SaveOperation(operation *domainMath.MathOperation, result *domainMath.Result) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.operations = append(r.operations, operation)
	r.results = append(r.results, result)
	return nil
}

// GetHistory returns all stored operations
func (r *InMemoryMathRepository) GetHistory() ([]*domainMath.MathOperation, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	// Return a copy to prevent external modifications
	history := make([]*domainMath.MathOperation, len(r.operations))
	copy(history, r.operations)
	return history, nil
}

// ClearHistory clears all stored operations
func (r *InMemoryMathRepository) ClearHistory() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.operations = make([]*domainMath.MathOperation, 0)
	r.results = make([]*domainMath.Result, 0)
	return nil
}
