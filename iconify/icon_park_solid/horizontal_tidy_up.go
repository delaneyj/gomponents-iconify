package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalTidyUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHorizontalTidyUp0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 8h6v32H7zm14 0h6v32h-6zm14 0h6v32h-6z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHorizontalTidyUp0)"/>`),
		g.Group(children),
	)
}