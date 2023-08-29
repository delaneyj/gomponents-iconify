package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeControl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12ZM8 8l3 3m1 11a9.99 9.99 0 0 0 8.307-4.43A9.953 9.953 0 0 0 22 12c0-5.523-4.477-10-10-10S2 6.477 2 12"/>`),
		g.Group(children),
	)
}