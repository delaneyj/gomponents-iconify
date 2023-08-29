package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 18V6l8.5 6l-8.5 6Zm-9 0V6l8.5 6L4 18Z"/>`),
		g.Group(children),
	)
}