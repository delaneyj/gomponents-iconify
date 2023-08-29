package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTwoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M5 22V2"/><path d="m5 14l2.47-.494a8.677 8.677 0 0 1 4.925.452a8.677 8.677 0 0 0 5.327.361l.1-.025a.9.9 0 0 0 .553-1.335l-1.56-2.601c-.342-.57-.513-.854-.553-1.163a1.508 1.508 0 0 1 0-.39c.04-.309.211-.594.553-1.163l1.278-2.13a.73.73 0 0 0-.803-1.084a7.3 7.3 0 0 1-4.482-.305l-.413-.165a8.677 8.677 0 0 0-4.924-.452L5 4" opacity=".5"/></g>`),
		g.Group(children),
	)
}