package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphicStitchingFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGraphicStitchingFour0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M39 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10ZM9 44a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm5-40H4v10h10V4Zm30 30H34v10h10V34Z"/><path d="M34 9H14m20 30H14m-5-5V14m30 20V14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGraphicStitchingFour0)"/>`),
		g.Group(children),
	)
}