package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Connection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSConnection0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="m39 34l5 5l-5 5"/><path fill="#fff" d="M8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="M12 8h8a4 4 0 0 1 4 4v23a4 4 0 0 0 4 4h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSConnection0)"/>`),
		g.Group(children),
	)
}