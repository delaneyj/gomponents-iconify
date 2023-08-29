package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 2h8.414l2 2h8.014l-2.364 6.5l2.364 6.5h-8.842l-2-2H5v7.5H3V2Zm2 11h6.414l2 2h5.158L17.3 11.5h-4.714l-2-2H5V13Zm0-5.5h6.414l2 2H17.3L18.572 6h-5.986l-2-2H5v3.5Z"/>`),
		g.Group(children),
	)
}