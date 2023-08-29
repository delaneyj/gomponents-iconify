package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Manjaro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2v20h5.6V7.6h7.2V2H2m7.2 7.2V22h5.6V9.2H9.2M16.4 2v20H22V2h-5.6Z"/>`),
		g.Group(children),
	)
}