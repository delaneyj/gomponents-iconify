package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Help(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 23a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0-5v-3c0-2 1-2 3-3s3-2.842 3-5A6 6 0 1 0 6 7"/>`),
		g.Group(children),
	)
}