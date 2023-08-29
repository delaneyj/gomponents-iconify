package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 5v43H128V5h128zm-85 278V155h42v128h-42zm171-141q42 52 42 119q0 80-56 136t-136 56t-136-56T0 261.5t56-136T192 69q67 0 120 43l30-31q16 13 30 30zM192 411q62 0 105.5-44T341 261t-43.5-105.5T192 112T86.5 155.5T43 261t43.5 106T192 411z"/>`),
		g.Group(children),
	)
}