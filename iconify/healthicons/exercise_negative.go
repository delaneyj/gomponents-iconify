package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExerciseNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsExerciseNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM36 14h-4v9H16v-9h-4v20h4v-9h16v9h4V14ZM6 17v6H4v2h2v6h4V17H6Zm36 8h2v-2h-2v-6h-4v14h4v-6Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsExerciseNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}