package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLamp0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" d="M14 9.5a3.5 3.5 0 1 1 7 0V21h-7V9.5Zm13 0a3.5 3.5 0 1 1 7 0V21h-7V9.5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 34h12v8H18z"/><path d="M10 22a1 1 0 0 1 1-1h26a1 1 0 0 1 1 1v4a8 8 0 0 1-8 8H18a8 8 0 0 1-8-8v-4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLamp0)"/>`),
		g.Group(children),
	)
}