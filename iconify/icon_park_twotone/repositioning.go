package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repositioning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRepositioning0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M33 22c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0Z"/><path fill="#555" d="M24 25a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="M9 14s7.5-11.5 20.5-7S42 24.5 42 24.5M42 8v16m-3 10s-6 11-19.5 7.5S6 24 6 24m0 0v16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRepositioning0)"/>`),
		g.Group(children),
	)
}