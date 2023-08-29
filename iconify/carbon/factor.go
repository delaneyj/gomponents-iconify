package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Factor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 28V9h-3V4h2V2h-6v2h2v5h-3v19h-4V16h-3v-5h2V9H8v2h2v5H7v12H4V2H2v26a2 2 0 0 0 2 2h26v-2Zm-14 0H9V18h4Zm8-17h4v17h-4Z"/>`),
		g.Group(children),
	)
}