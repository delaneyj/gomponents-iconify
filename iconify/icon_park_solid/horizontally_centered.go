package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontallyCentered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHorizontallyCentered0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke-linejoin="round" d="M16 16h16v16H16z"/><path d="M5 40V8m38 32V8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHorizontallyCentered0)"/>`),
		g.Group(children),
	)
}