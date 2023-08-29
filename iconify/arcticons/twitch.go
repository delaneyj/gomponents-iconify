package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.16 12.16v8.19h2.47v-8.19Zm-7.75 0v8.19H26v-8.19Zm-9.3-7.66l-6.88 6.84v25.12h8.3v7l7.06-7h5.61L40.77 24V4.5Zm1.42 2.89H38v15.2L32.55 28h-5.61l-5.12 5.13V28h-6.29Z"/>`),
		g.Group(children),
	)
}