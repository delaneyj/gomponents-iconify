package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BombEmojiBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M10 16h2"/><path fill="currentColor" d="M14 12.5c0 .828-.448 1.5-1 1.5s-1-.672-1-1.5s.448-1.5 1-1.5s1 .672 1 1.5Z"/><ellipse cx="9" cy="12.5" fill="currentColor" rx="1" ry="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m17 7l-2 2"/><path stroke="currentColor" stroke-width="1.5" d="M17.981 2.353a.558.558 0 0 1 1.038 0l.654 1.66c.057.143.17.257.315.314l1.659.654c.47.186.47.852 0 1.038l-1.66.654a.558.558 0 0 0-.314.315l-.654 1.659a.558.558 0 0 1-1.038 0l-.654-1.66a.558.558 0 0 0-.315-.314l-1.659-.654a.558.558 0 0 1 0-1.038l1.66-.654a.558.558 0 0 0 .314-.315l.654-1.659Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M5.75 8.003a7.5 7.5 0 1 1-2.747 2.747"/></g>`),
		g.Group(children),
	)
}