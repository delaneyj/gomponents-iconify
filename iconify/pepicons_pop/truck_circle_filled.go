package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TruckCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopTruckCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M10.5 21.75a2.25 2.25 0 1 1 0-4.5a2.25 2.25 0 0 1 0 4.5Zm9 0a2.25 2.25 0 1 1 0-4.5a2.25 2.25 0 0 1 0 4.5Z"/><path fill-rule="evenodd" d="M21 7.5H11a2 2 0 0 0-2 2V17a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V9.5a2 2 0 0 0-2-2ZM11 17V9.5h10V17H11Z" clip-rule="evenodd"/><path d="M13.25 11.75a.5.5 0 0 1 0-1h5.5a.5.5 0 0 1 0 1h-5.5Zm0 2a.5.5 0 0 1 0-1h5.5a.5.5 0 0 1 0 1h-5.5Zm0 2a.5.5 0 0 1 0-1h5.5a.5.5 0 0 1 0 1h-5.5Z"/><path fill-rule="evenodd" d="M9 10.5H5.736a2 2 0 0 0-1.92 1.442L3 14.75V17a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4.5a2 2 0 0 0-2-2ZM5 17v-1.966l.736-2.534H9V17H5Z" clip-rule="evenodd"/><path d="M6.29 13.1H8a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5H6a.5.5 0 0 1-.5-.5v-.85l.015-.122l.29-1.15a.5.5 0 0 1 .485-.378Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopTruckCircleFilled0)"/></g>`),
		g.Group(children),
	)
}