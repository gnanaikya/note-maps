// Code generated by "kvschema"; DO NOT EDIT.

package storage

import (
	"github.com/google/note-maps/kv"
)

// Store provides entities, components, and indexes backed by a key-value
// store.
//
// Usage:
//
//   t, err := Store{Store: store}.TopicMapInfoComponent(0).Scan([]kv.Entity{7, 42})
//
type Store struct {
	kv.Store
	parent kv.Entity
}

func (s Store) Parent(e kv.Entity) *Store {
	s.parent = e
	return &s
}

// SetTopicMapInfo sets the TopicMapInfo associated with e to v.
//
// Corresponding indexes are updated.
func (s *Store) SetTopicMapInfo(e kv.Entity, v TopicMapInfo) error {
	key := make(kv.Prefix, 8+2+8)
	s.parent.EncodeAt(key)
	TopicMapInfoPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	return s.Set(key, v.Encode())
}

// GetTopicMapInfoSlice returns a TopicMapInfo for each entity in es.
//
// If the underlying storage returns an empty value with no error for keys that
// do not exist, and TopicMapInfo.Decode() can decode an empty byte slice, then a
// query for entities that are not associated with a TopicMapInfo should return no
// errors.
func (s *Store) GetTopicMapInfoSlice(es []kv.Entity) ([]TopicMapInfo, error) {
	result := make([]TopicMapInfo, len(es))
	key := make(kv.Prefix, 8+2+8)
	s.parent.EncodeAt(key)
	TopicMapInfoPrefix.EncodeAt(key[8:])
	for i, e := range es {
		e.EncodeAt(key[10:])
		err := s.Get(key, (&result[i]).Decode)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}