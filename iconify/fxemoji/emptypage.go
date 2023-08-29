package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emptypage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#F9E7C0" d="M437.567 512H88.004a8.182 8.182 0 0 1-8.182-8.182V8.182A8.182 8.182 0 0 1 88.004 0H437.83a7.92 7.92 0 0 1 7.92 7.92v495.898a8.183 8.183 0 0 1-8.183 8.182z"/>`),
		g.Group(children),
	)
}