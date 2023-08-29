package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewportWide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2H2v4h2V4h16v2h2V2H4zM2 20v-2h2v2h16v-2h2v4H2v-2zm16-9h-5v2h5v2h-2v2h2v-2h2v-2h2v-2h-2V9h-2V7h-2v2h2v2zm-7 0H6V9h2V7H6v2H4v2H2v2h2v2h2v2h2v-2H6v-2h5v-2z"/>`),
		g.Group(children),
	)
}