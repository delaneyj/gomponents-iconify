package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTvOne0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M42 12H6a2 2 0 0 0-2 2v26a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2V14a2 2 0 0 0-2-2Z"/><path fill="#000" stroke="#000" d="M31 19H11v16h20V19Z"/><path stroke="#fff" stroke-linecap="round" d="m14 4.5l9.09 7.5L34 2"/><path stroke="#000" stroke-linecap="round" d="M38 18v1m0 6v1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTvOne0)"/>`),
		g.Group(children),
	)
}