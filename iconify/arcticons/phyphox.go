package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phyphox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.283 39.328l2.942-.414M21.231 7.043L20.963 5m-.379 3.538l-4.354.633a3.987 3.987 0 0 0-3.392 4.502l3.64 25.896a3.987 3.987 0 0 0 4.5 3.392l10.424-1.465a3.987 3.987 0 0 0 3.393-4.5l-3.64-25.897a3.988 3.988 0 0 0-4.502-3.393l-4.408.608m-6.122 28.947l18.317-2.574M5.736 26.65l2.481 3.736l.943-9.321l2.613 4.737l1.22-.172"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.929 25.217l1.942-.272l1.652-9.349l5.504 13.746l2.131-9.432l2.334 3.4l1.79-.25m3.195-.45l2.913-.41l.697-4.076l2.716 5.425l1.461-2.033"/>`),
		g.Group(children),
	)
}