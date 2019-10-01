//
// Copyright (c) 2017 The heketi Authors
//
// This file is licensed to you under your choice of the GNU Lesser
// General Public License, version 3 or any later version (LGPLv3 or
// later), or the GNU General Public License, version 2 (GPLv2), in all
// cases as published by the Free Software Foundation.
//

package glusterfs

import (
	"bytes"
	"encoding/gob"

	"github.com/boltdb/bolt"
	"github.com/lpabon/godbc"
)

var (
	dbAttributeKeys = []string{
		DB_BRICK_HAS_SUBTYPE_FIELD,
		DB_CLUSTER_HAS_FILE_BLOCK_FLAG,
		DB_GENERATION_ID,
		DB_HAS_PENDING_OPS_BUCKET,
	}
)

type DbAttributeEntry struct {
	Key   string
	Value string
}

func NewDbAttributeEntry() *DbAttributeEntry {
	entry := &DbAttributeEntry{}
	return entry
}

func NewDbAttributeEntryFromKey(tx *bolt.Tx, key string) (*DbAttributeEntry, error) {

	entry := NewDbAttributeEntry()
	err := EntryLoad(tx, entry, key)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (dba *DbAttributeEntry) BucketName() string {
	return BOLTDB_BUCKET_DBATTRIBUTE
}

func (dba *DbAttributeEntry) Save(tx *bolt.Tx) error {
	godbc.Require(tx != nil)
	godbc.Require(len(dba.Key) > 0)

	return EntrySave(tx, dba, dba.Key)
}

func (dba *DbAttributeEntry) Delete(tx *bolt.Tx) error {
	godbc.Require(tx != nil)

	return EntryDelete(tx, dba, dba.Key)
}

func (dba *DbAttributeEntry) Marshal() ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(*dba)

	return buffer.Bytes(), err
}

func (dba *DbAttributeEntry) Unmarshal(buffer []byte) error {
	dec := gob.NewDecoder(bytes.NewReader(buffer))
	err := dec.Decode(dba)
	if err != nil {
		return err
	}

	return nil
}

func DbAttributeList(tx *bolt.Tx) ([]string, error) {
	list := EntryKeys(tx, BOLTDB_BUCKET_DBATTRIBUTE)
	if list == nil {
		return nil, ErrAccessList
	}
	return list, nil
}

// validDbAttributeKeys returns true if all dbattribute keys in the
// database match keys in knownKeys map.
func validDbAttributeKeys(tx *bolt.Tx, knownKeys map[string]bool) bool {
	list := EntryKeys(tx, BOLTDB_BUCKET_DBATTRIBUTE)
	if list == nil {
		logger.LogError("unable to list keys in dbattribute bucket")
		return false
	}
	for _, key := range list {
		if !knownKeys[key] {
			logger.LogError("unknown dbattribute key: %+v", key)
			return false
		}
	}
	return true
}

// mapDbAtrributeKeys returns a map equivalent of dbAttributeKeys
// for fast lookup.
func mapDbAtrributeKeys() map[string]bool {
	m := map[string]bool{}
	for _, k := range dbAttributeKeys {
		m[k] = true
	}
	return m
}
