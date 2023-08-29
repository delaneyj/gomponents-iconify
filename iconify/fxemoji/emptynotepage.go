package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emptynotepage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#FFD469" d="M444.027 452.829H67.973a8.802 8.802 0 0 1-8.802-8.802V67.973a8.802 8.802 0 0 1 8.802-8.802h376.215a8.64 8.64 0 0 1 8.64 8.64v376.215a8.801 8.801 0 0 1-8.801 8.803z"/>`),
		g.Group(children),
	)
}