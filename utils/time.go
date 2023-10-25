package utils

import "time"

//ParseTime : func to parse various times in string format
func ParseTime(dateString string) (time.Time, error) {
	var parsedDate time.Time
	var err error
	parsedDate, err = time.Parse("2006-01-02T15:04:05.000Z", dateString)
	if err != nil {
		parsedDate, err = time.Parse("2006-01-02T15:04:05.000-07:00", dateString)
		if err != nil {
			parsedDate, err = time.Parse("2006-01-02T15:04:05Z", dateString)
			if err != nil {
				parsedDate, err = time.Parse("2018-10-18T19:40:15+02:00", dateString)
				if err != nil {
					parsedDate, err = time.Parse("2018-02-22 07:25:03.063+01", dateString)
					if err != nil {
						parsedDate, err = time.Parse("2006-01-02T03:04:05.999-07:00", dateString)
						if err != nil {
							parsedDate, err = time.Parse("2006-01-02T03:04:05.999999", dateString)
							if err != nil {
								parsedDate, err = time.Parse("2006-01-02T03:04:05-07:00", dateString)
								if err != nil {
									parsedDate, err = time.Parse("2006-01-02T03:04:05.9+02:00", dateString)
									if err != nil {
										parsedDate, err = time.Parse("2006-01-02T15:04:05.99+02:00", dateString)
										if err != nil {
											parsedDate, err = time.Parse("2006-01-02 15:04:05.999999", dateString)
											if err != nil {
												parsedDate, err = time.Parse("2006-01-02", dateString)
												if err != nil {
													parsedDate, err = time.Parse("2006-01-02 15:04", dateString)
													if err != nil {
														return time.Time{}, err
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return parsedDate, nil
}
