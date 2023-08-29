package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Googlelens(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16.667a4.666 4.666 0 1 0 0-9.333a4.666 4.666 0 0 0 0 9.333m8 6a2.666 2.666 0 1 0 0-5.333a2.666 2.666 0 0 0 0 5.333m-13.333-2a3.343 3.343 0 0 1-3.334-3.334v-2.666H0v2.666A6.665 6.665 0 0 0 6.667 24h2.666v-3.333zm-3.334-14c0-1.834 1.5-3.334 3.334-3.334h2.666V0H6.667A6.665 6.665 0 0 0 0 6.667v2.666h3.333zm14-3.334c1.834 0 3.334 1.5 3.334 3.334v2.666H24V6.667A6.665 6.665 0 0 0 17.333 0h-2.666v3.333Z"/>`),
		g.Group(children),
	)
}