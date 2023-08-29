package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autoscaling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 28H12v-2h10V10H6v10H4V10a2.002 2.002 0 0 1 2-2h16a2.002 2.002 0 0 1 2 2v16a2.002 2.002 0 0 1-2 2Z"/><path fill="currentColor" d="M28 22h-2v-2h2V4H12v2h-2V4a2.002 2.002 0 0 1 2-2h16a2.002 2.002 0 0 1 2 2v16a2.002 2.002 0 0 1-2 2Z"/><path fill="currentColor" d="M10 14v2h4.586L3 27.586L4.414 29L16 17.414V22h2v-8h-8z"/>`),
		g.Group(children),
	)
}