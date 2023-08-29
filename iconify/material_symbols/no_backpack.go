package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoBackpack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.5 13.675l3.5 3.5V8q0-1.4-.85-2.45T17 4.15V2h-3v2h-4V2H7v2.15L14.825 12H16.5v1.675Zm3.275 8.925l-3.2-3.2h2.8L20 20q0 .825-.588 1.413T18 22H6q-.825 0-1.413-.588T4 20V8q0-.825.313-1.55t.862-1.275l.975.975v2.825L1.4 4.225L2.8 2.8l18.375 18.4l-1.4 1.4ZM7.5 14h3.675l-2-2H7.5v2Z"/>`),
		g.Group(children),
	)
}