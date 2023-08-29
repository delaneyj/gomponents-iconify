package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M11.75 12.416h48V60.25h-48z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 12.166h48v48H12z"/><path d="M31.34 40.518h0a4.346 4.346 0 0 0 4.345 4.346h.63a4.346 4.346 0 0 0 4.346-4.346h0a4.346 4.346 0 0 0-4.346-4.346h-.63a4.346 4.346 0 0 0-4.346 4.346Zm0-8.704h0a4.346 4.346 0 0 0 4.345 4.346h.63a4.346 4.346 0 0 0 4.346-4.346h0a4.346 4.346 0 0 0-4.346-4.346h-.63a4.346 4.346 0 0 0-4.346 4.346Z"/></g>`),
		g.Group(children),
	)
}