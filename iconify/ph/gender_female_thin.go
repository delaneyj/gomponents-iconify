package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderFemaleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 96a76 76 0 1 0-80 75.89V204H88a4 4 0 0 0 0 8h36v28a4 4 0 0 0 8 0v-28h36a4 4 0 0 0 0-8h-36v-32.11A76.09 76.09 0 0 0 204 96ZM60 96a68 68 0 1 1 68 68a68.07 68.07 0 0 1-68-68Z"/>`),
		g.Group(children),
	)
}