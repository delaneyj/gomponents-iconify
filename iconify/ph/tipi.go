package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tipi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M238.74 211.69L137.5 53.5l21.24-33.19a8 8 0 0 0-13.48-8.62L128 38.66l-17.26-27a8 8 0 1 0-13.48 8.62L118.5 53.5L17.26 211.69A8 8 0 0 0 24 224h208a8 8 0 0 0 6.74-12.31ZM86.3 208l41.7-65.16L169.7 208Zm102.4 0l-54-84.31a8 8 0 0 0-13.48 0L67.3 208H38.62L128 68.34L217.38 208Z"/>`),
		g.Group(children),
	)
}