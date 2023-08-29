package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScoopAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.806 8.089h24.388V42.5H11.806zm-4.97 10.733h34.328M6.836 35.776h34.328M6.836 27.3h34.328M14.875 5.5l3.926 3.759M33.125 5.5l-3.926 3.758"/>`),
		g.Group(children),
	)
}