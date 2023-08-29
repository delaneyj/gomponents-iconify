package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M2.998 1H17.5L21 4.5V23H3L2.998 1ZM16 1v5h5m-9 14v-9m-4 3l4-4l4 4"/>`),
		g.Group(children),
	)
}