package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarcodeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6v12m7-12v12m7-12v12M10 6.5v11a.5.5 0 0 0 1 0v-11a.5.5 0 0 0-1 0zm7 0v11a.5.5 0 0 0 1 0v-11a.5.5 0 0 0-1 0zM3.5 6v0a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .5.5v0a.5.5 0 0 0 .5-.5v-11a.5.5 0 0 0-.5-.5z"/>`),
		g.Group(children),
	)
}