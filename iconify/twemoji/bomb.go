package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bomb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#CCD6DD" d="m24.187 9.657l5.658-5.654L32 6.16l-5.658 5.655z"/><circle cx="14" cy="22" r="14" fill="#31373D"/><path fill="#31373D" d="m19 11.342l5.658-5.657l5.657 5.658L24.657 17z"/><circle cx="32" cy="4" r="4" fill="#F18F26"/><circle cx="32" cy="4" r="2" fill="#FDCB58"/>`),
		g.Group(children),
	)
}