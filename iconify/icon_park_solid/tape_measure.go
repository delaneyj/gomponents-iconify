package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TapeMeasure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTapeMeasure0"><g fill="none" stroke-width="4"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M13 15v-3a3 3 0 0 0-3-3v0a3 3 0 0 0-3 3v7"/><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M4 27c0-6.627 5.373-12 12-12s12 5.373 12 12v12H4V27Zm24 0h16v12H28z"/><circle cx="16" cy="27" r="4" fill="#000" stroke="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M39 27v4m-6-4v4"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M30 27h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTapeMeasure0)"/>`),
		g.Group(children),
	)
}