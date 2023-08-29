package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Haxe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M0 0h32.3L64 16L96.3 0H128v32.6L111.8 64L128 95.8V128H95.2L64 112.1L32.6 128H0V94.7L15.7 64L0 31.8z"/>`),
		g.Group(children),
	)
}