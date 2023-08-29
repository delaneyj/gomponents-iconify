package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRectangle0"><path fill="#555" stroke="#fff" stroke-width="4" d="M38 4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRectangle0)"/>`),
		g.Group(children),
	)
}