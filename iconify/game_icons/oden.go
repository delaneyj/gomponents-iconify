package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M480.2 31.81L351.5 50.26L461.7 160.5zm-92.7 79.99l-13.1 13.1c2.3 1.8 4.6 3.8 6.7 5.9c2.2 2.2 4.2 4.5 6 6.9l13.2-13.2zm-58.1 15.5c-14 0-28.1 5.4-38.9 16.2c-21.5 21.5-21.5 56.2 0 77.8c21.6 21.5 56.3 21.5 77.8 0c21.6-21.6 21.6-56.3 0-77.8c-10.8-10.8-24.8-16.2-38.9-16.2zm-127.1 82l-32.6 32.5l100.5 100.5l32.5-32.6zm69.6 18l-13.1 13.1l12.8 12.8l13.1-13.1c-2.4-1.9-4.7-3.9-6.9-6.1c-2.1-2.1-4.1-4.4-5.9-6.7zm-71 71.1l-169.09 169l12.72 12.8L213.6 311.1z"/>`),
		g.Group(children),
	)
}