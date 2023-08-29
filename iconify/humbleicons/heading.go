package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heading(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4v8m0 8v-8m10-8v8m0 8v-8m0 0H7M5 4h4m6 0h4m0 16h-4m-6 0H5"/>`),
		g.Group(children),
	)
}