package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarouselVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="26" height="34" x="11" y="7" fill="#2F88FF" stroke="#000"/><rect width="7" height="26" x="4" y="11" stroke="#000"/><rect width="7" height="26" x="37" y="11" stroke="#000"/><path fill="#43CCF8" stroke="#fff" d="M22 20L28 24L22 28V20Z"/></g>`),
		g.Group(children),
	)
}