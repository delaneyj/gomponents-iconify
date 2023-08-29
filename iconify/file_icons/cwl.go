package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cwl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m65.642 272.597l83.288-83.235l-78.548-75.911l78.185-79.468L115.249 0L2.185 114.282l79.904 75.774L0 272.49l80.261 80.732L0 435.093L75.366 512l32.983-33.641l-42.618-43.668l80.313-81.286z"/>`),
		g.Group(children),
	)
}