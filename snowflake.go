// Snowflake is snowflake as defined https://en.wikipedia.org/wiki/Snowflake_ID
// But we use the most significant bit, and a signed timestamp
// its format is:
//
//     42 bits     10 bits     12 bits
//  | 64 ... 23 | 22 ... 12 |  11 ... 0  |
//	| timestamp | machineID | sequenceID |
//
package snowflake

import "time"

var seq int64

// NewSafe creates a new snowflake ID checking for overflows
// given a specifed epoch and machine ID
func NewSafe(epoch time.Time, mID int) int64 {

	// time in milliseconds since our epoch
	s := time.Now().Unix() - epoch.Unix()

	if s >= 1<<42 {
		return 0
	}

	if mID >= 1<<10 {
		return 0
	}

	seq = (seq + 1) % (1 << 12)

	// arithmetic shift
	return (s << 22) | (int64(mID) << 12) | (seq)
}

// New creates a new snowflake ID without checking for overflow
func New(epoch time.Time, mID int) int64 {
	seq = (seq + 1) % (1 << 12)
	s := time.Now().Sub(epoch).Milliseconds()
	return (s << 22) | (int64(mID) << 12) | (seq)
}

// Unpack gets the timestamp, the machine ID and the sequence ID from a
func Unpack(s int64, epoch time.Time) (stamp time.Time, machineID, sequenceID int64) {
	t := time.UnixMilli(epoch.UnixMilli() + s>>22)
	return t, (s >> 12) & ((1 << 10) - 1), s & ((1 << 13) - 1)
}
