package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpMonitorHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.11 12.45L14 10.24l-3.11 6.21c-.16.34-.51.55-.89.55s-.73-.21-.89-.55L7.38 13H2v7h20v-7h-6c-.38 0-.73-.21-.89-.55z"/><path fill="currentColor" d="M22 4H2v7h6c.38 0 .73.21.89.55L10 13.76l3.11-6.21a1 1 0 0 1 1.79 0L16.62 11H22V4z"/>`),
		g.Group(children),
	)
}