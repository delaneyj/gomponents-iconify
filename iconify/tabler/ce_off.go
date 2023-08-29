package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 4a7.99 7.99 0 0 0-2.581.426M5.867 5.864A8 8 0 0 0 11 20m9-16a8 8 0 0 0-7.29 4.7M12 12a8 8 0 0 0 8 8m-4-8h4M3 3l18 18"/>`),
		g.Group(children),
	)
}