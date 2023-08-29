package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 21H7c-1.654 0-3-1.346-3-3V6c0-1.654 1.346-3 3-3h10c1.654 0 3 1.346 3 3v12c0 1.654-1.346 3-3 3zM7 5c-.551 0-1 .449-1 1v12c0 .551.449 1 1 1h10c.551 0 1-.449 1-1V6c0-.551-.449-1-1-1H7zm9 6H8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1zm0-3H8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1zm0 6H8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1zm0 3H8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1z"/>`),
		g.Group(children),
	)
}