package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorI(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorI0" d="M36 21a1 1 0 0 0-1 1v28a1 1 0 1 0 2 0V22a1 1 0 0 0-1-1Z"/></defs><use href="#openmojiRegionalIndicatorI0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorI0"/></g>`),
		g.Group(children),
	)
}