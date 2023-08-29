package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Okrs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-7a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM3 5h10M3 12h10M3 19h10m3 2.243l2.121-2.122m0 0L20.243 17m-2.122 2.121L16 17m2.121 2.121l2.122 2.122"/>`),
		g.Group(children),
	)
}