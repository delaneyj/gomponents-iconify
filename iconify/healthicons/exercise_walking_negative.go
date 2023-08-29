package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExerciseWalkingNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsExerciseWalkingNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM25.693 28.397l5.193 5.124a2 2 0 0 1 .457.693l2.769 7.055a2 2 0 1 1-3.724 1.462l-2.614-6.661l-8.928-8.81a2 2 0 0 1-.583-1.649l.715-6.32c-1.724 1.714-3.054 4.123-4.073 7.316a2 2 0 0 1-3.81-1.216c1.87-5.86 4.975-10.246 10.185-12.257l.023-.009c1.327-.493 2.707-.453 3.937.182c1.181.611 2.022 1.666 2.573 2.848c.232.498.446.963.648 1.4c.488 1.058.898 1.95 1.293 2.732c.553 1.1.998 1.83 1.438 2.341c.408.475.813.767 1.33.969c.556.217 1.335.367 2.538.403a2 2 0 0 1-.12 3.998c-1.445-.043-2.728-.228-3.873-.675c-1.183-.462-2.116-1.165-2.91-2.09c-.5-.582-.94-1.247-1.35-1.97l-1.114 5.134Zm-7.43 1.823l3.315 3.18l-1.526 5.147a2 2 0 0 1-.684 1.006l-5.135 4.023a2 2 0 0 1-2.466-3.15l4.632-3.628l1.395-4.71l.469-1.868ZM31.25 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsExerciseWalkingNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}