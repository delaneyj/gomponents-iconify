package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServiceInstance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21h2.03v-6H3v6zm9.45 1l-2.93-.64l-2.93-.64A.76.76 0 0 1 6 20v-4a.71.71 0 0 1 .17-.47a.75.75 0 0 1 .43-.27l4.78-.59l4.79-.67l.15.55l.15.55a1 1 0 0 1-.56.9l-2 .5l-2 .5l1.62.65l1.62.65l2.28-.64l2.22-.66a.73.73 0 0 1 .58.07a.75.75 0 0 1 .36.47l.19.76l.2.77a.74.74 0 0 1-.54.9l-3.63 1l-3.63 1a1.49 1.49 0 0 1-.36.05a1.55 1.55 0 0 1-.37-.02Zm5.57-13.5l-3.5-2.25L11.02 4v9l3.5-2.25l3.5-2.25z"/>`),
		g.Group(children),
	)
}