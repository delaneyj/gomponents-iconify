package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperShip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke-linecap="round" d="M4 24L12.5714 42L24 29L4 24Z"/><path fill="#2F88FF" stroke-linecap="round" d="M44 24L35.4286 42L24 29L44 24Z"/><path fill="#2F88FF" stroke-linecap="round" d="M13 42L35 42L24 29L13 42Z"/><path d="M12 26L24 4L36 26"/></g>`),
		g.Group(children),
	)
}