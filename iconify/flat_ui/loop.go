package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<circle cx="50" cy="50" r="50" fill="#E74C3C"/><path fill="none" stroke="#F0C419" stroke-dasharray="216 0" stroke-dashoffset="43" stroke-linecap="square" stroke-miterlimit="10" stroke-width="8" d="M12.5 50.5c0 10 7 17 17 17c17 0 26-34 43-34c10 0 18 7 18 17s-8 17-18 17c-17 0-26-34-43-34c-10 0-17 7-17 17z"/><path d="m45.938 50.5l5.047 6.516s-2.141 2.797-3.5 4.25c0 0-3.621-3.552-5.746-5.583c1.308-1.402 4.199-5.183 4.199-5.183z" opacity=".15"/>`),
		g.Group(children),
	)
}