package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Balance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBalance0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path fill="#fff" stroke="#fff" d="M35 25a32.234 32.234 0 0 0-22 0l-1-11c7-3 17-3 24 0l-1 11Z"/><path stroke="#000" d="m24 23l-3-5"/><path stroke="#fff" d="M42 39a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v30Z"/><path stroke="#fff" d="M29 23.455a32.222 32.222 0 0 0-10 0"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBalance0)"/>`),
		g.Group(children),
	)
}