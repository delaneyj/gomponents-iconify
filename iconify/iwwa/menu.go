package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M32.666 13.25H7.334a.5.5 0 0 1 0-1h25.332a.5.5 0 0 1 0 1zm0 7.25H7.334a.5.5 0 0 1 0-1h25.332a.5.5 0 0 1 0 1zm0 7.25H7.334a.5.5 0 0 1 0-1h25.332a.5.5 0 0 1 0 1z"/>`),
		g.Group(children),
	)
}