package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCup0"><g fill="none" stroke="#fff" stroke-width="4"><path d="M8.778 17.012c0-.559.453-1.012 1.012-1.012h23.976c.559 0 1.012.453 1.012 1.012V31c0 7.18-5.82 13-13 13s-13-5.82-13-13V17.012Z"/><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="M8.778 23h26v8h-26z"/><path stroke-linecap="round" stroke-linejoin="round" d="M21.778 4v6m-8-4v2m16-2v2"/><path stroke-linecap="round" d="M34.778 34a7 7 0 1 0 0-14"/></g></mask><path fill="currentColor" d="M0 0h49v48H0z" mask="url(#ipTCup0)"/>`),
		g.Group(children),
	)
}