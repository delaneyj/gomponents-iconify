package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Poll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h18v18H7V7zm8 3v12h2V10h-2zm-5 4v8h2v-8h-2zm10 2v6h2v-6h-2z"/>`),
		g.Group(children),
	)
}