package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataMining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.403 8.69a7.11 7.11 0 0 0-3.348-4.217l-.869-.496l.496-.868l-1.736-.993l-.496.869L4.977 1l-.992 1.736l3.473 1.985l-.497.868l-.992 1.737L2 14.272l1.736.992l3.97-6.946l.992-1.736l.496-.869l.868.496Z"/><path fill="currentColor" d="M10 10v3h12v-3Zm11 2h-4v-1h4Zm-11 3v3h12v-3Zm11 2h-4v-1h4Zm-11 3v3h12v-3Zm11 2h-4v-1h4Z"/>`),
		g.Group(children),
	)
}