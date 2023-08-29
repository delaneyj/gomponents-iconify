package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfacePageControllerFitScreenFitScreenAdjustDisplayArtboardFrameCorner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 .5H1a.5.5 0 0 0-.5.5v4m13 0V1a.5.5 0 0 0-.5-.5H9m0 13h4a.5.5 0 0 0 .5-.5V9M.5 9v4a.5.5 0 0 0 .5.5h4"/>`),
		g.Group(children),
	)
}