package course

// getCoursesLevel returns a slice of strings representing the levels of courses available.
// The levels are "Beginner", "Intermediate", and "Advanced".
func GetCoursesLevel() []string {
	return []string{
		"Beginner",
		"Intermediate",
		"Advanced",
	}
}

// getCoursesByLevel returns a slice of strings representing the courses available for a given level.
// The level parameter should be one of "Beginner", "Intermediate", or "Advanced".
// The function returns a slice of course names for the specified level.
// If the level is not recognized, it returns an empty slice.
// The courses are hardcoded for now, but can be replaced with a database or API call in the future.
func GetCoursesByLevel(level string) []string {
	courses := map[string][]string{
		"Beginner": {
			"golang-101: Beginner Golang",
			"golang-102: Beginner Golang",
		},
		"Intermediate": {
			"golang-201: Intermediate Golang",
			"golang-202: Intermediate Golang",
		},
		"Advanced": {
			"golang-301: Advanced Golang",
			"golang-302: Advanced Golang",
			"golang-303: Advanced Golang",
		},
	}

	return courses[level]
}

// GetCourseByCourseID returns a string representing the course name for a given course ID.
// The courseID parameter should be a string representing the course ID.
// The function returns the course name for the specified course ID.
// If the course ID is not recognized, it returns an empty string.
// The courses are hardcoded for now, but can be replaced with a database or API call in the future.
func GetCourseByCourseID(courseID string) string {
	courses := map[string]string{
		"golang-101": "Beginner Golang - 101",
		"golang-102": "Beginner Golang - 102",
		"golang-201": "Intermediate Golang - 201",
		"golang-202": "Intermediate Golang - 202",
		"golang-301": "Advanced Golang - 301",
		"golang-302": "Advanced Golang - 302",
		"golang-303": "Advanced Golang - 303",
	}

	return courses[courseID]
}
