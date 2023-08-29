package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 6v20h24V6zm2 2h5v4H6zm7 0h13v4H13zm-7 6h5v4H6zm7 0h13v4H13zm-7 6h5v4H6zm7 0h13v4H13z"/>`),
		g.Group(children),
	)
}