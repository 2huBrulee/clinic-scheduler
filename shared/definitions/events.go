package definitions

type eventName string

const (
	DoctorCreated eventName = "doctor-created"
)

var eventNames = []eventName{
	DoctorCreated,
}

func eventNamesToString(eventNames []eventName) []string {
	names := make([]string, 0)

	for _, name := range eventNames {
		names = append(names, string(name))
	}

	return names
}

var StringEventNames = eventNamesToString(eventNames)
