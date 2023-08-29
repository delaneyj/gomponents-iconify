package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buzzfeed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M32 16c0 8.839-7.161 16-16 16S0 24.839 0 16S7.161 0 16 0s16 7.161 16 16zm-5.531-.364L25.166 6.38l-8.667 3.5l3.432 1.984l-3.244 5.62l-5.62-3.244l-5.537 9.588l3.1 1.792l3.749-6.489l5.62 3.244l5.036-8.724z"/>`),
		g.Group(children),
	)
}