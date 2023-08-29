package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M6 16h10H6Zm0-4h12H6Zm0-4h5h-5Zm8-7v7h7M3 23V1h12l6 6v16H3Z"/>`),
		g.Group(children),
	)
}