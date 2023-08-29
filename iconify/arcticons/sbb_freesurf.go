package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SbbFreesurf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.87 6.041l2.865 2.868l-2.867 2.867m-5.739 0L18.265 8.91l2.867-2.868m-2.867 2.867h11.47M24 6.041v5.736"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.421 15.318h39.158M25.848 33.592a3.68 3.68 0 1 0-3.696 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.797 36.982a7.59 7.59 0 1 0-7.645-.03"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.774 40.357a11.5 11.5 0 1 0-11.548 0"/><circle cx="24" cy="30.409" r="1.249" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 31.657v10.252"/>`),
		g.Group(children),
	)
}