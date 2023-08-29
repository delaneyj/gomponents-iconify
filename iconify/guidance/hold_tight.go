package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HoldTight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M2 8.5h4.5A3.5 3.5 0 0 0 10 5h4.5m3 0h3v.5a3 3 0 0 1-3 3h-4h5.25a1.75 1.75 0 1 1 0 3.5H16.5h1.25a1.75 1.75 0 1 1 0 3.5H16.5h.25a1.75 1.75 0 1 1 0 3.5H2M16 8.5V0m0 24v-5"/>`),
		g.Group(children),
	)
}