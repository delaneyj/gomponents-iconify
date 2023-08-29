package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fsharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#378BBA" d="M0 121.492L121.492 0v60.746l-60.746 60.746l60.746 60.745v60.746L0 121.492Z"/><path fill="#378BBA" d="m78.102 121.492l43.39-43.39v86.78l-43.39-43.39Z"/><path fill="#30B9DB" d="M256 121.492L130.17 0v60.746l60.745 60.746l-60.746 60.745v60.746L256 121.492Z"/>`),
		g.Group(children),
	)
}