package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laugh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M358 0c198 0 359 155 359 349S556 698 358 698S0 543 0 349S160 0 358 0zM138 290l3 1c5 0 10-3 12-8c0-1 11-49 55-49c45 0 55 47 56 49c2 5 6 8 11 8l32-10c3-1 5-3 6-6c2-3 2-6 1-8c-1-4-20-85-106-85c-91 0-105 81-106 84s-1 6 1 9c1 3 4 5 7 6zm437 1l32-10c3-1 5-3 6-6c2-3 2-6 1-8c-1-4-20-85-105-85c-91 0-106 81-107 84s-1 6 1 9c1 3 4 5 7 6l28 9l3 1c5 0 10-3 12-8c0-1 11-49 56-49c44 0 55 47 55 49c2 5 6 8 11 8zm-21 157c7-20-4-43-24-51s-325-8-344 0c-20 8-30 31-23 51c32 81 108 134 195 134s164-53 196-134z"/>`),
		g.Group(children),
	)
}