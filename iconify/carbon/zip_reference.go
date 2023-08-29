package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZipReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4zM28 2h-6v14h2v-5h4a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm0 7h-4V4h4zM12 4h3v10h-3v2h8v-2h-3V4h3V2h-8v2zM2 4h6L2 14v2h8v-2H4l6-10V2H2v2z"/>`),
		g.Group(children),
	)
}