package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.518 1.232a.556.556 0 0 0-.495.268l-4 6.66c-.222.37.045.84.477.84h8c.432 0 .7-.47.477-.84l-4-6.66a.556.556 0 0 0-.46-.268zm.002.922L8.431 7H7.76L6.416 5.773L5.519 7l-.894-1.227L3.281 7H2.61l2.91-4.846z" fill="currentColor"/>`),
		g.Group(children),
	)
}