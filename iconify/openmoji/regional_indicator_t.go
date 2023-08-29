package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorT(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorT0" d="M44 21H28a1 1 0 1 0 0 2h7v27a1 1 0 1 0 2 0V23h7a1 1 0 1 0 0-2Z"/></defs><use href="#openmojiRegionalIndicatorT0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorT0"/></g>`),
		g.Group(children),
	)
}