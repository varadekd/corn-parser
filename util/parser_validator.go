package util

import "regexp"

func ValidCornParser(cron string) bool {
	cronRegex := `^(\*|([0-9]|[1-5][0-9])|(\*\/([1-9]|[1-5][0-9]))|(([0-9]|[1-5][0-9])\-([0-9]|[1-5][0-9])))(,(\*|([0-9]|[1-5][0-9])|(\*\/([1-9]|[1-5][0-9]))|(([0-9]|[1-5][0-9])\-([0-9]|[1-5][0-9]))))* ` +
		`(\*|([0-9]|1[0-9]|2[0-3])|(\*\/([1-9]|1[0-9]|2[0-3]))|(([0-9]|1[0-9]|2[0-3])\-([0-9]|1[0-9]|2[0-3])))(,(\*|([0-9]|1[0-9]|2[0-3])|(\*\/([1-9]|1[0-9]|2[0-3]))|(([0-9]|1[0-9]|2[0-3])\-([0-9]|1[0-9]|2[0-3]))))* ` +
		`(\*|([1-9]|[12][0-9]|3[01])|(\*\/([1-9]|[12][0-9]|3[01]))|(([1-9]|[12][0-9]|3[01])\-([1-9]|[12][0-9]|3[01])))(,(\*|([1-9]|[12][0-9]|3[01])|(\*\/([1-9]|[12][0-9]|3[01]))|(([1-9]|[12][0-9]|3[01])\-([1-9]|[12][0-9]|3[01]))))* ` +
		`(\*|([1-9]|1[0-2])|(\*\/([1-9]|1[0-2]))|(([1-9]|1[0-2])\-([1-9]|1[0-2])))(,(\*|([1-9]|1[0-2])|(\*\/([1-9]|1[0-2]))|(([1-9]|1[0-2])\-([1-9]|1[0-2]))))* ` +
		`(\*|([0-6])|(\*\/([0-6]))|(([0-6])\-([0-6])))(,(\*|([0-6])|(\*\/([0-6]))|(([0-6])\-([0-6]))))* .+$`

	match, _ := regexp.MatchString(cronRegex, cron)
	return match
}
