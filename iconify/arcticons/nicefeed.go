package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nicefeed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="11.371" cy="21.99" r="5.871" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="rotate(-.073 11.382 22.014)"/><circle cx="22.277" cy="11.371" r="5.871" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.386 11.37a26.015 26.015 0 0 1-26.015 26.016"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 11.37A31.129 31.129 0 0 1 11.371 42.5"/><path fill="none" stroke="currentColor" d="M37.386 11.37H42.5M11.371 37.387V42.5"/>`),
		g.Group(children),
	)
}