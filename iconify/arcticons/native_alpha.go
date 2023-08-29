package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NativeAlpha(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 9.335a5.5 5.5 0 0 0-3.921 0c-12.23 5.332-9.17 29.594-22.509 29.33C10.633 38.535 4.504 33.097 4.5 24S10.633 9.463 17.07 9.335c13.34-.265 10.279 23.997 22.508 29.33a5.5 5.5 0 0 0 3.922 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.74 29.5v-11l8 11v-11"/>`),
		g.Group(children),
	)
}