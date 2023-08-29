package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNewComputer0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="36" height="28" x="6" y="6" fill="#fff" rx="3"/><path stroke-linecap="round" d="M14 42h20m-10-8v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNewComputer0)"/>`),
		g.Group(children),
	)
}