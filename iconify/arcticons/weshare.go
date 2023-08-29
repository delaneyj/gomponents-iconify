package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weshare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.998 2.5A21.51 21.51 0 0 0 5.152 34.35L2.504 45.487l11.138-2.647A21.498 21.498 0 1 0 23.998 2.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.522 20.228l5.401 13.061a26.32 26.32 0 0 1 11.275-2.639a23.94 23.94 0 0 1 10.938 2.64V20.227a12.87 12.87 0 0 0-4.456-1.76l-2.768-6.97a25.377 25.377 0 0 0-11.005 2.166c-7.832 3.451-9.385 6.564-9.385 6.564Z"/>`),
		g.Group(children),
	)
}