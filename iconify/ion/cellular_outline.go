package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellularOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<rect width="64" height="320" x="416" y="96" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="8" ry="8"/><rect width="64" height="240" x="288" y="176" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="8" ry="8"/><rect width="64" height="176" x="160" y="240" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="8" ry="8"/><rect width="64" height="112" x="32" y="304" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="8" ry="8"/>`),
		g.Group(children),
	)
}