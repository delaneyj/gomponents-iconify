package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 5h-3m-5 0h5m-5 0v10m0-10h-1c-1.667 0-5 1-5 5s3.333 5 5 5h1m0 4v-4m5 4V5"/>`),
		g.Group(children),
	)
}