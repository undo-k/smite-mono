package main

// type Duration struct {
// 	time.Duration
// }

// func (d Duration) MarshalJSON() ([]byte, error) {
// 	minutes := d.Duration / time.Minute
// 	seconds := (d.Duration % time.Minute) / time.Second
// 	return json.Marshal(fmt.Sprintf("%02d:%02d", minutes, seconds))
// }

// func (d *Duration) UnmarshalJSON(b []byte) error {
// 	var v interface{}
// 	if err := json.Unmarshal(b, &v); err != nil {
// 		return err
// 	}
// 	switch value := v.(type) {
// 	case float64:
// 		d.Duration = time.Duration(value)
// 		return nil
// 	case string:
// 		var err error
// 		d.Duration, err = time.ParseDuration(value)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	default:
// 		return errors.New("invalid duration")
// 	}
// }
