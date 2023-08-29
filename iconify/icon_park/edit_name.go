package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditName(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="11" r="7" fill="#2F88FF"/><path d="M4 41C4 32.1634 12.0589 25 22 25"/><path fill="#2F88FF" d="M31 42L41 32L37 28L27 38V42H31Z"/></g>`),
		g.Group(children),
	)
}