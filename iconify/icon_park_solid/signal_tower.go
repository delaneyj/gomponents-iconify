package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalTower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSignalTower0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M17.081 6c-6.12 4.285-7.607 12.72-3.322 18.84c4.285 6.12 12.72 7.608 18.84 3.323L17.082 6Z"/><path d="M22 31v11m-9-17.5L6 42h36M40 6L25 17M17 6h23l-7 21.5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSignalTower0)"/>`),
		g.Group(children),
	)
}