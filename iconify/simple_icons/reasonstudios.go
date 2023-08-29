package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reasonstudios(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.49 5.114l8.3-4.79a2.421 2.421 0 0 1 2.39-.017l.03.017l8.299 4.79a2.42 2.42 0 0 1 1.211 2.065v9.611a2.42 2.42 0 0 1-1.184 2.08l-.027.016l-8.299 4.79a2.42 2.42 0 0 1-2.39.017l-.03-.017l-8.3-4.79a2.421 2.421 0 0 1-1.21-2.065V7.21c0-.855.45-1.645 1.184-2.08l.026-.016l8.3-4.79zM12 4.026L5.092 8.013v7.974L12 19.974V12l6.908-3.987z"/>`),
		g.Group(children),
	)
}