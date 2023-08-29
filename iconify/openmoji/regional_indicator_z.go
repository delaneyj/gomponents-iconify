package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorZ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorZ0" d="M44 49H29.724l15.144-26.504A1 1 0 0 0 44 21H28a1 1 0 1 0 0 2h14.276L27.132 49.504A1 1 0 0 0 28 51h16a1 1 0 1 0 0-2Z"/></defs><use href="#openmojiRegionalIndicatorZ0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorZ0"/></g>`),
		g.Group(children),
	)
}