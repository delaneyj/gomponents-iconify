package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CitySix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 2H6v6H1v14h22V6h-5V2Zm0 6h3v12h-3V8Zm-2 12h-3v-5h-2v5H8V4h8v16ZM6 20H3V10h3v10Z"/>`),
		g.Group(children),
	)
}