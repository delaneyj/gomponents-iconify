package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTEyes0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" fill-rule="evenodd" stroke-linejoin="round" d="M24 41c9.941 0 18-8.322 18-14c0-5.678-8.059-14-18-14S6 21.328 6 27c0 5.672 8.059 14 18 14Z" clip-rule="evenodd"/><path fill="#555" stroke-linejoin="round" d="M24 33a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path stroke-linecap="round" d="m13.264 11.266l2.594 3.62m19.767-3.176l-2.595 3.62M24.009 7v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTEyes0)"/>`),
		g.Group(children),
	)
}