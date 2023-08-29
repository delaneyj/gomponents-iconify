package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asciicircum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 516L230 11l6-11h68l6 11l231 505h-79L270 97L79 516H0z"/>`),
		g.Group(children),
	)
}