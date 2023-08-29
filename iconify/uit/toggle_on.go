package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 6.5h-9a5.5 5.5 0 0 0 0 11h9a5.5 5.5 0 0 0 0-11zm0 10h-9a4.5 4.5 0 1 1 0-9h9a4.5 4.5 0 1 1 0 9zm0-7a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5zm0 4a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3z"/>`),
		g.Group(children),
	)
}