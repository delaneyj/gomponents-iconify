package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.477 2 2 6.483 2 12.014c0 5.278 4.078 9.602 9.25 9.986v-6.061a3.505 3.505 0 0 1-2.75-3.424v-.501a1 1 0 0 1 1-1.002h.25V9.01a.75.75 0 1 1 1.5 0v2.002h1.5V9.01a.75.75 0 1 1 1.5 0v2.002h.25a1 1 0 0 1 1 1.002v.5a3.505 3.505 0 0 1-2.75 3.425V22c5.172-.384 9.25-4.708 9.25-9.986C22 6.484 17.523 2 12 2Z"/>`),
		g.Group(children),
	)
}