package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="32" height="28" x="4" y="10" fill="#2F88FF" stroke="#000"/><path stroke="#000" stroke-linecap="round" d="M44 14L36 20.75V27.25L44 34V14Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" d="M17 19L23 24L17 29"/></g>`),
		g.Group(children),
	)
}