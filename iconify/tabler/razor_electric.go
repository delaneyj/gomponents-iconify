package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RazorElectric(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3v2m4-2v2m4-2v2m-7 7v6a3 3 0 0 0 6 0v-6H9zM8 5h8l-1 4H9zm4 12v1"/>`),
		g.Group(children),
	)
}