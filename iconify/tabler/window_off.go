package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.166 6.19A6.903 6.903 0 0 0 5 10v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-1m0-4v-5c0-3.728-3.134-7-7-7a6.86 6.86 0 0 0-3.804 1.158M5 13h8m4 0h2M12 3v5m0 4v9M3 3l18 18"/>`),
		g.Group(children),
	)
}