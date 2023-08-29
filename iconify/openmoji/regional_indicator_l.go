package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorL(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorL0" d="M43 49.012H32v-27a1 1 0 1 0-2 0v28a1 1 0 0 0 1 1h12a1 1 0 1 0 0-2Z"/></defs><use href="#openmojiRegionalIndicatorL0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorL0"/></g>`),
		g.Group(children),
	)
}