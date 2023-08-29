package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustifySpaceAround(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 22V2h2v20h-2ZM2 22V2h2v20H2Zm13-5V7h3v10h-3Zm-9 0V7h3v10H6Z"/>`),
		g.Group(children),
	)
}