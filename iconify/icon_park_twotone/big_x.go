package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BigX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBigX0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M33 6h11L15 42H4L33 6Z"/><path d="M15 6H4l29 36h11L15 6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBigX0)"/>`),
		g.Group(children),
	)
}