package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestDesktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M18.218 1H23v18H1V1h5m11 9c-4-3-6 3-10 0M5 23h14H5Zm5-22v4.773l-5 7.182V15h14v-2.045l-5-7.182V1M8 1h8h-8Zm0 22h8v-4H8v4Z"/>`),
		g.Group(children),
	)
}