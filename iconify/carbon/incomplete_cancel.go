package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncompleteCancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 21.41L28.59 20L25 23.59L21.41 20L20 21.41L23.59 25L20 28.59L21.41 30L25 26.41L28.59 30L30 28.59L26.41 25L30 21.41zM14 26a12 12 0 0 1 0-24zm3.826-21.236a10.029 10.029 0 0 1 3.242 2.168L22.48 5.52a12.036 12.036 0 0 0-3.89-2.602zM26 14a11.93 11.93 0 0 0-.917-4.59l-1.847.764A9.943 9.943 0 0 1 24 14z"/>`),
		g.Group(children),
	)
}