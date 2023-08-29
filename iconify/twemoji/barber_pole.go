package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarberPole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="6" r="6" fill="#CCD6DD"/><path fill="#FFF" d="M11 12h14v21H11z"/><path fill="#DD2E44" d="M11 28.487L20.251 33H25v-2.134l-14-6.83z"/><path fill="#55ACEE" d="m11 19.585l14 6.83v-4.45l-14-6.831z"/><path fill="#DD2E44" d="M13.697 12L25 17.514V12z"/><path fill="#99AAB5" d="M27 11a2 2 0 0 1-2 2H11a2 2 0 0 1 0-4h14a2 2 0 0 1 2 2zm0 23a2 2 0 0 1-2 2H11a2 2 0 0 1 0-4h14a2 2 0 0 1 2 2z"/>`),
		g.Group(children),
	)
}