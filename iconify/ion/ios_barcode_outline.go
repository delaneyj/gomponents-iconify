package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosBarcodeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M48 384h80v-16H64V144h64v-16H48z" fill="currentColor"/><path d="M384 128v16h64v224h-64v16h80V128z" fill="currentColor"/><path d="M112 192h16v128h-16z" fill="currentColor"/><path d="M384 192h16v128h-16z" fill="currentColor"/><path d="M320 160h16v192h-16z" fill="currentColor"/><path d="M176 160h16v192h-16z" fill="currentColor"/><path d="M247 176h16v160h-16z" fill="currentColor"/>`),
		g.Group(children),
	)
}