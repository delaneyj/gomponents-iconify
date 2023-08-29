package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spreaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.615 0l-5.64 6.54L.529 4.718l8.68 7.372l-8.537 7.463l8.411-1.984L14.843 24l.71-8.601l7.918-3.483l-7.963-3.33L14.621 0h-.006z"/>`),
		g.Group(children),
	)
}