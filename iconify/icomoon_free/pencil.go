package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.5 0a2.5 2.5 0 0 1 2 4l-1 1L11 1.5l1-1c.418-.314.937-.5 1.5-.5zM1 11.5L0 16l4.5-1l9.25-9.25l-3.5-3.5L1 11.5zm10.181-5.819l-7 7l-.862-.862l7-7l.862.862z"/>`),
		g.Group(children),
	)
}