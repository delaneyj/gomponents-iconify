package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExerciseRunningNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsExerciseRunningNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM23.337 31.651l-4.242-5.52a2 2 0 0 1-.282-1.934l2.647-6.915c-1.247.31-2.098.776-2.785 1.354c-1.03.868-1.862 2.12-2.986 4.088a2 2 0 1 1-3.473-1.985c1.119-1.957 2.25-3.788 3.883-5.162c1.704-1.436 3.79-2.248 6.603-2.561c1.176-.13 2.47-.104 3.683.437c1.273.569 2.203 1.588 2.837 2.95c.854 1.833 1.489 2.924 1.997 3.556c.245.304.416.445.513.508c.077.05.11.054.122.056h.001c.087.01.369 0 1.197-.367c.361-.16.755-.352 1.237-.587l.115-.056a53.62 53.62 0 0 1 1.784-.84a2 2 0 0 1 1.625 3.655a49.38 49.38 0 0 0-1.653.779l-.131.064c-.461.225-.925.452-1.36.644c-.9.398-2.05.83-3.297.679c-1.316-.16-2.33-.903-3.165-1.9l-2.733 5.273l3.78 4.918c.225.294.364.644.403 1.012l.832 7.996a2 2 0 1 1-3.978.414l-.774-7.433l-2.296-2.988l-.02.037l-.084-.172Zm-6.377 1.317l1.472-3.96l2.876 3.742l-1.125 3.028a2 2 0 0 1-2.033 1.297l-7.308-.581a2 2 0 0 1 .317-3.988l5.801.462ZM35 8.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsExerciseRunningNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}