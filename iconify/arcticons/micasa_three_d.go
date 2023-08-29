package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicasaThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.87 26.12a4.419 4.419 0 0 1 4.418-4.418h0a4.419 4.419 0 0 1 4.42 4.419v7.291m-8.839-11.71v11.71"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.707 26.12a4.419 4.419 0 0 1 4.42-4.418h0a4.419 4.419 0 0 1 4.418 4.419v7.291"/><circle cx="33.234" cy="18.484" r="3.896" fill="none" stroke="currentColor" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.234 26.247v7.165"/>`),
		g.Group(children),
	)
}