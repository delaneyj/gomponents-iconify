package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yoga(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m2 14.5l1.654-.551a6 6 0 0 0 3.795-3.795L8 8.5s1.5-1 4-1s4 1 4 1l.551 1.654a6 6 0 0 0 3.795 3.795L22 14.5m-12.5-4v5l-4.25.85a.937.937 0 0 0-.042 1.825c1.305.332 4.198 1.095 6.792 1.99c2.657.916 5 1.97 5 2.835m-2.5-12.5v5l4.25.85a.937.937 0 0 1 .042 1.825c-.974.248-2.833.736-4.792 1.34m-4 1.397c-1.73.712-3 1.45-3 2.088m4.85-17.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}