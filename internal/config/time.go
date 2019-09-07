package config

import "time"

// Duration is custom type for config
type Duration time.Duration

// UnmarshalText unmarshal text to time.Duration
func (d *Duration) UnmarshalText(t []byte) error {
	dur, err := time.ParseDuration(string(t))
	if err != nil {
		return err
	}
	*d = Duration(dur)
	return nil

}
