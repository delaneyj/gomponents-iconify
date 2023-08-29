package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 17c-2.727 0-6-2.778-6-5s3.273-5 6-5s6 2.778 6 5s-3.273 5-6 5Zm-1-5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm9-7L4 19"/>`),
		g.Group(children),
	)
}