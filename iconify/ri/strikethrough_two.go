package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrikethroughTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 9h-2V6H5V4h14v2h-6v3Zm0 6v5h-2v-5h2ZM3 11h18v2H3v-2Z"/>`),
		g.Group(children),
	)
}