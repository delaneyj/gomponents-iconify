package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileGifFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 2l5 5v13.992A1 1 0 0 1 20.007 22H3.993A1 1 0 0 1 3 21.008V2.992C3 2.444 3.447 2 3.998 2H16Zm-3 8h-1v5h1v-5Zm-2 0H9a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h1a1 1 0 0 0 1-1v-2H9v1h1v1H9a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1h2v-1Zm6 0h-3v5h1v-2h2v-1h-2v-1h2v-1Z"/>`),
		g.Group(children),
	)
}