package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileZipFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 2v2h2V2h8.007c.548 0 .993.444.993.992v18.016a1 1 0 0 1-.993.992H3.993A.993.993 0 0 1 3 21.008V2.992A1 1 0 0 1 3.993 2H10Zm2 2v2h2V4h-2Zm-2 2v2h2V6h-2Zm2 2v2h2V8h-2Zm-2 2v2h2v-2h-2Zm2 2v2h-2v3h4v-5h-2Z"/>`),
		g.Group(children),
	)
}