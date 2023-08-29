package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorC0" d="M36.682 23c2.204 0 4.326.804 5.973 2.264a1 1 0 1 0 1.327-1.496a10.995 10.995 0 0 0-7.3-2.768c-6.065 0-11 4.935-11 11v8c0 6.065 4.935 11 11 11c2.694 0 5.287-.983 7.3-2.768a1 1 0 1 0-1.327-1.496A8.995 8.995 0 0 1 36.682 49c-4.963 0-9-4.038-9-9v-8c0-4.962 4.038-9 9-9Z"/></defs><use href="#openmojiRegionalIndicatorC0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorC0"/></g>`),
		g.Group(children),
	)
}