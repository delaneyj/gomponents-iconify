package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackLimitation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 22h8v8H8zm16-10h-4v2h4v2h-3v2h3v2h-4v2h4a2.003 2.003 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2zm-8 8H8v-8h8zm-6-2h4v-4h-4zm6-14.59L14.59 2L12 4.59L9.41 2L8 3.41L10.59 6L8 8.59L9.41 10L12 7.41L14.59 10L16 8.59L13.41 6L16 3.41z"/>`),
		g.Group(children),
	)
}