package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorA0" d="M36.919 21.606a1 1 0 0 0-1.838 0l-12 28a1 1 0 1 0 1.838.788l2.62-6.114h16.922l2.62 6.114a1 1 0 1 0 1.838-.788l-12-28ZM28.396 42.28L36 24.538l7.604 17.742H28.396Z"/></defs><use href="#openmojiRegionalIndicatorA0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorA0"/></g>`),
		g.Group(children),
	)
}