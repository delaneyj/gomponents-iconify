package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorJ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorJ0" d="M44.346 21a1 1 0 0 0-1 1v18c0 4.962-4.037 9-9 9a9.002 9.002 0 0 1-6.024-2.312a1.001 1.001 0 0 0-1.336 1.488A11 11 0 0 0 34.346 51c6.066 0 11-4.935 11-11V22a1 1 0 0 0-1-1Z"/></defs><use href="#openmojiRegionalIndicatorJ0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorJ0"/></g>`),
		g.Group(children),
	)
}