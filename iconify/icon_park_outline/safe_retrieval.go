package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafeRetrieval(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M7 9.127L24.008 4L41 9.127v10.515C41 30.695 34.153 40.506 24.003 44C13.849 40.506 7 30.692 7 19.638V9.128Z"/><path d="M24 30a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path stroke-linecap="round" d="m29 29l6 7"/><path d="M41 19.643C41 30.695 34.153 40.507 24.002 44"/></g>`),
		g.Group(children),
	)
}