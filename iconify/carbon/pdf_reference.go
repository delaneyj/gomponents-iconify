package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PdfReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4zm18-4h2v-6h5V8h-5V4h6V2h-8v14zM16 2h-4v14h4a4 4 0 0 0 4-4V6a4 4 0 0 0-4-4zm2 10a2 2 0 0 1-2 2h-2V4h2a2 2 0 0 1 2 2zM8 2H2v14h2v-5h4a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm0 7H4V4h4z"/>`),
		g.Group(children),
	)
}