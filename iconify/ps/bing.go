package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M2 0v246q2 64 69 109t161 45q95 0 162.5-46.5T462 242t-67.5-111.5T232 84q-82 0-147 36V0H2zm230 137q61 0 104 31t43 74t-43 74t-104 31t-104-31t-43-74t43-74t104-31zm-97 105q0-29 28.5-49t68.5-20t68.5 20t28.5 49t-28.5 49.5T232 312t-68.5-20.5T135 242z"/>`),
		g.Group(children),
	)
}