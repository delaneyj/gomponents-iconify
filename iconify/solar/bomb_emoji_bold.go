package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BombEmojiBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M17.981 2.353a.558.558 0 0 1 1.038 0l.654 1.66c.057.143.17.257.315.314l1.659.654c.47.186.47.852 0 1.038l-1.66.654a.558.558 0 0 0-.314.315l-.654 1.659a.558.558 0 0 1-1.038 0l-.654-1.66a.558.558 0 0 0-.315-.314l-1.659-.654a.558.558 0 0 1 0-1.038l1.66-.654a.558.558 0 0 0 .314-.315l.654-1.659Z"/><path fill-rule="evenodd" d="M17 14.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0Zm-5 2.25a.75.75 0 0 0 0-1.5h-2a.75.75 0 0 0 0 1.5h2Zm2-4.25c0 .828-.448 1.5-1 1.5s-1-.672-1-1.5s.448-1.5 1-1.5s1 .672 1 1.5ZM9 14c.552 0 1-.672 1-1.5S9.552 11 9 11s-1 .672-1 1.5s.448 1.5 1 1.5Z" clip-rule="evenodd"/><path d="m16.767 8.294l-.75.75a8.555 8.555 0 0 0-1.06-1.061l.75-.75l.76.3l.3.76Z"/></g>`),
		g.Group(children),
	)
}