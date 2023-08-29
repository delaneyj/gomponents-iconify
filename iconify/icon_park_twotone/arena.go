package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arena(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTArena0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M44 26L24 36L4 26l20-10l20 10Z"/><path d="m24 7l20 10l-20 10L4 17L24 7Z"/><path fill="#555" stroke-linecap="round" d="M44 26v8L24 44L4 34v-8l20 10l20-10Z"/><path stroke-linecap="round" d="M44 14v12M4 26V14m20 22V24m0-8V4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTArena0)"/>`),
		g.Group(children),
	)
}