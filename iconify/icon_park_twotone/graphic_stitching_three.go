package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphicStitchingThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGraphicStitchingThree0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm0 30a5 5 0 1 0 0-10a5 5 0 0 0 0 10ZM14 19H4v10h10V19Zm30 0H34v10h10V19Z"/><path d="M19 9H9v10m10 20H9V29M29 9h11v10M29 39h10V29"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGraphicStitchingThree0)"/>`),
		g.Group(children),
	)
}