package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InvertCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTInvertCamera0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M9 14s7.5-11.5 20.5-7S42 24.5 42 24.5M39 34s-6 11-19.5 7.5S6 24 6 24M42 8v16M6 24v16"/><path fill="#555" stroke-linecap="round" d="M14 20h12v8H14z"/><path d="m34 28l-2-1.333v-5.334L34 20v8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTInvertCamera0)"/>`),
		g.Group(children),
	)
}