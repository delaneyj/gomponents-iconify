package skill_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProcessingDark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="none"><rect width="256" height="256" fill="#242938" rx="60"/><path stroke="#0468FF" stroke-width="42.286" d="M138.572 170.762c84.571 0 84.571-112.762 0-112.762"/><path stroke="#1F34AB" stroke-width="42.286" d="M138.571 86.19L54 198.952"/><path stroke="#85AEFF" stroke-width="42.286" d="m54 114.381l28.19 56.381"/></g>`),
		g.Group(children),
	)
}