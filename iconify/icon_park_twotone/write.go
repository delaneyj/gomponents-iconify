package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Write(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWrite0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M5.325 43.5h8.485l31.113-31.113l-8.486-8.485L5.325 35.015V43.5Z"/><path stroke-linecap="round" d="m27.952 12.387l8.485 8.485"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWrite0)"/>`),
		g.Group(children),
	)
}