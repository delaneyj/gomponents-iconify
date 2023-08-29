package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagMore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M469 0q18 0 30.5 12.5T512 43v298q0 18-12.5 30.5T469 384H151q-23 0-36-19L0 192L115 19q13-19 34-19h320zM192 224q13 0 22.5-9.5T224 192t-9.5-22.5T192 160t-22.5 9.5T160 192t9.5 22.5T192 224zm106.5 0q13.5 0 23-9.5T331 192t-9.5-22.5t-23-9.5t-22.5 9.5t-9 22.5t9 22.5t22.5 9.5zm107 0q13.5 0 22.5-9.5t9-22.5t-9-22.5t-22.5-9.5t-23 9.5T373 192t9.5 22.5t23 9.5z"/>`),
		g.Group(children),
	)
}