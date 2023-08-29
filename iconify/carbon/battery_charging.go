package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryCharging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 11h-1v-1a2 2 0 0 0-2-2h-4v2h4v3h3v6h-3v3h-5v2h5a2 2 0 0 0 2-2v-1h1a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2zM11 22H6V10h6V8H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h5z"/><path fill="currentColor" d="m14.81 23.58l-1.62-1.16L17.06 17H9.37l6.85-8.62l1.56 1.24L13.51 15h7.43l-6.13 8.58z"/>`),
		g.Group(children),
	)
}