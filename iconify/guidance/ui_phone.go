package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UiPhone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M9.5.5V2A1.5 1.5 0 0 0 11 3.5h2A1.5 1.5 0 0 0 14.5 2V.5m-4.5 20h4M4.5.75S5 7.492 5 12s-.5 11.25-.5 11.25v.25h15v-.25S19 16.508 19 12S19.5.75 19.5.75V.5h-15v.25Z"/>`),
		g.Group(children),
	)
}