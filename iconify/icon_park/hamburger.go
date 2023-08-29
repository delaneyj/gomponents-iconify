package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hamburger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" d="M44 22C44 12.0589 35.0457 4 24 4C12.9543 4 4 12.0589 4 22H44Z" clip-rule="evenodd"/><rect width="40" height="6" x="4" y="38" fill="#2F88FF"/><path d="M4 28L9.45455 32L16.7273 28L24 32L31.2727 28L38.5455 32L44 28"/></g>`),
		g.Group(children),
	)
}