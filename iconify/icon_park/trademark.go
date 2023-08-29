package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trademark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#2F88FF" stroke="#000"/><path stroke="#fff" d="M12 19H16L20 19"/><path stroke="#fff" d="M16 19L16 29"/><path stroke="#fff" d="M26 29V19L31 25L36 19V29"/></g>`),
		g.Group(children),
	)
}