package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Accessibility(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.55 26.11l-3.05 1.52L23.66 21H15a2 2 0 0 1-2-2v-6a2 2 0 0 1 4 0v4h7v-2h-5v-2a4 4 0 0 0-8 0v1a9 9 0 1 0 8.77 11h-2.06A7 7 0 1 1 11 16v3a4 4 0 0 0 4 4h7.34l3.16 7.37l4.95-2.48zM15.5 8A3.5 3.5 0 1 1 19 4.5A3.5 3.5 0 0 1 15.5 8zm0-5A1.5 1.5 0 1 0 17 4.5A1.5 1.5 0 0 0 15.5 3z"/>`),
		g.Group(children),
	)
}