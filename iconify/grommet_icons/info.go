package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Info(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M15 17c0-3 4-5 4-9s-3-7-7-7s-7 3-7 7s4 6 4 9v3c0 2 1 3 3 3s3-1 3-3v-3Zm-6 1h6"/>`),
		g.Group(children),
	)
}