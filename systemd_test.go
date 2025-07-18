package systemdconf

import (
	"reflect"
	"testing"
	"time"
)

func TestParseBool(t *testing.T) {
	tests := []struct {
		String  string
		Success bool
		Boolean bool
	}{
		// true
		{
			String:  "1",
			Success: true,
			Boolean: true,
		},
		{
			String:  "y",
			Success: true,
			Boolean: true,
		},
		{
			String:  "Y",
			Success: true,
			Boolean: true,
		},
		{
			String:  "yes",
			Success: true,
			Boolean: true,
		},
		{
			String:  "YES",
			Success: true,
			Boolean: true,
		},
		{
			String:  "true",
			Success: true,
			Boolean: true,
		},
		{
			String:  "TRUE",
			Success: true,
			Boolean: true,
		},
		{
			String:  "on",
			Success: true,
			Boolean: true,
		},
		{
			String:  "ON",
			Success: true,
			Boolean: true,
		},

		// false
		{
			String:  "0",
			Success: true,
			Boolean: false,
		},
		{
			String:  "n",
			Success: true,
			Boolean: false,
		},
		{
			String:  "N",
			Success: true,
			Boolean: false,
		},
		{
			String:  "no",
			Success: true,
			Boolean: false,
		},
		{
			String:  "NO",
			Success: true,
			Boolean: false,
		},
		{
			String:  "false",
			Success: true,
			Boolean: false,
		},
		{
			String:  "FALSE",
			Success: true,
			Boolean: false,
		},
		{
			String:  "off",
			Success: true,
			Boolean: false,
		},
		{
			String:  "OFF",
			Success: true,
			Boolean: false,
		},

		// Error
		{
			String:  "garbage",
			Success: false,
		},
		{
			String:  "",
			Success: false,
		},
		{
			String:  "full",
			Success: false,
		},
	}
	for _, td := range tests {
		t.Run("String="+td.String+";ParseBool", func(t *testing.T) {
			b, err := ParseBool(td.String)
			if td.Success {
				if err != nil {
					t.Errorf("ParseBool() = _, %v; want nil", err)
				}
				if b != td.Boolean {
					t.Errorf("ParseBool() = %v, _; want %v", b, td.Boolean)
				}
			} else {
				if err != ErrInvalidBool {
					t.Errorf("ParseBool() = _, %v; want ErrInvalidBool", err)
				}
			}
		})
		t.Run("String="+td.String+";Value.Bool", func(t *testing.T) {
			v := Value{td.String}
			b, err := v.Bool()
			if td.Success {
				if err != nil {
					t.Errorf("Value.Bool() = _, %v; want nil", err)
				}
				if b != td.Boolean {
					t.Errorf("Value.Bool() = %v, _; want %v", b, td.Boolean)
				}
			} else {
				if err != ErrInvalidBool {
					t.Errorf("Value.Bool() = _, %v; want ErrInvalidBool", err)
				}
			}
		})
		t.Run("String="+td.String+";Value.BoolSlice", func(t *testing.T) {
			v := Value{"true", td.String}
			b, err := v.BoolSlice()
			if td.Success {
				if err != nil {
					t.Errorf("Value.BoolSlice() = _, %v; want nil", err)
				}
				expected := []bool{true, td.Boolean}
				if !reflect.DeepEqual(b, expected) {
					t.Errorf("Value.BoolSlice() = %v, _; want %v", b, expected)
				}
			} else {
				if err != ErrInvalidBool {
					t.Errorf("Value.BoolSlice() = _, %v; want ErrInvalidBool", err)
				}
			}
		})
	}
}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		String   string
		Success  bool
		Duration time.Duration
	}{
		// OK
		{
			String:   "5s",
			Success:  true,
			Duration: 5 * time.Second,
		},
		{
			String:   "5s500ms",
			Success:  true,
			Duration: 5*time.Second + 500*time.Millisecond,
		},
		{
			String:   " 5s 500ms  ",
			Success:  true,
			Duration: 5*time.Second + 500*time.Millisecond,
		},
		{
			String:   " 5.5s  ",
			Success:  true,
			Duration: 5*time.Second + 500*time.Millisecond,
		},
		{
			String:   " 5.5s 0.5ms ",
			Success:  true,
			Duration: 5*time.Second + 500*time.Millisecond + 500*time.Microsecond,
		},
		{
			String:   " .22s ",
			Success:  true,
			Duration: 220 * time.Millisecond,
		},
		{
			String:   " .50y ",
			Success:  true,
			Duration: year / 2,
		},
		{
			String:   "2.5",
			Success:  true,
			Duration: 2500 * time.Millisecond,
		},
		{
			String:   ".7",
			Success:  true,
			Duration: 700 * time.Millisecond,
		},
		{
			String:   "23us",
			Success:  true,
			Duration: 23 * time.Microsecond,
		},
		{
			String:   "23µs",
			Success:  true,
			Duration: 23 * time.Microsecond,
		},
		{
			String:   "infinity",
			Success:  true,
			Duration: infinity,
		},
		{
			String:   " infinity ",
			Success:  true,
			Duration: infinity,
		},
		{
			String:   "+3.1s",
			Success:  true,
			Duration: 3100 * time.Millisecond,
		},
		{
			String:   "3.1s.2",
			Success:  true,
			Duration: 3300 * time.Millisecond,
		},
		{
			String:   "3.1 .2",
			Success:  true,
			Duration: 3300 * time.Millisecond,
		},
		{
			String:   "3.1 sec .2 sec",
			Success:  true,
			Duration: 3300 * time.Millisecond,
		},
		{
			String:   "3.1 sec 1.2 sec",
			Success:  true,
			Duration: 4300 * time.Millisecond,
		},

		// Error
		{
			String:  " xyz ",
			Success: false,
		},
		{
			String:  "",
			Success: false,
		},
		{
			String:  " . ",
			Success: false,
		},
		{
			String:  " 5. ",
			Success: false,
		},
		{
			String:  ".s ",
			Success: false,
		},
		{
			String:  "-5s ",
			Success: false,
		},
		{
			String:  "-0.3s ",
			Success: false,
		},
		{
			String:  "-0.0s ",
			Success: false,
		},
		{
			String:  "-0.-0s ",
			Success: false,
		},
		{
			String:  "0.-0s ",
			Success: false,
		},
		{
			String:  "3.-0s ",
			Success: false,
		},
		{
			String:  " infinity .7",
			Success: false,
		},
		{
			String:  ".3 infinity",
			Success: false,
		},
		{
			String:  "3.+1s",
			Success: false,
		},
		{
			String:  "3. 1s",
			Success: false,
		},
		{
			String:  "3.s",
			Success: false,
		},
		{
			String:  "12.34.56",
			Success: false,
		},
		{
			String:  "12..34",
			Success: false,
		},
		{
			String:  "..1234",
			Success: false,
		},
		{
			String:  "1234..",
			Success: false,
		},
	}
	for _, td := range tests {
		t.Run("String="+td.String+";ParseDuration", func(t *testing.T) {
			d, err := ParseDuration(td.String, time.Second)
			if td.Success {
				if err != nil {
					t.Errorf("ParseDuration() = _, %v; want nil", err)
				}
				if d != td.Duration {
					t.Errorf("ParseDuration() = %v, _; want %v", d, td.Duration)
				}
			} else {
				if err != ErrInvalidDuration {
					t.Errorf("ParseDuration() = _, %v; want ErrInvalidDuration", err)
				}
			}
		})
		t.Run("String="+td.String+";Value.Duration", func(t *testing.T) {
			v := Value{td.String}
			d, err := v.Duration()
			if td.Success {
				if err != nil {
					t.Errorf("Value.Duration() = _, %v; want nil", err)
				}
				if d != td.Duration {
					t.Errorf("Value.Duration() = %v, _; want %v", d, td.Duration)
				}
			} else {
				if err != ErrInvalidDuration {
					t.Errorf("Value.Duration() = _, %v; want ErrInvalidDuration", err)
				}
			}
		})
		t.Run("String="+td.String+";Value.DurationSlice", func(t *testing.T) {
			v := Value{"3s", td.String}
			d, err := v.DurationSlice()
			if td.Success {
				if err != nil {
					t.Errorf("Value.DurationSlice() = _, %v; want nil", err)
				}
				expected := []time.Duration{3 * time.Second, td.Duration}
				if !reflect.DeepEqual(d, expected) {
					t.Errorf("Value.DurationSlice() = %v, _; want %v", d, expected)
				}
			} else {
				if err != ErrInvalidDuration {
					t.Errorf("Value.DurationSlice() = _, %v; want ErrInvalidDuration", err)
				}
			}
		})
	}
}
