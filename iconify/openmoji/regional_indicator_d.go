package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorD0" d="M38.694 21.052h-9.708a1 1 0 0 0-1 1v27.896a1 1 0 0 0 1 1h9.708c5.139 0 9.32-4.18 9.32-9.32v-11.26c0-5.137-4.181-9.316-9.32-9.316Zm7.32 20.576c0 4.036-3.284 7.32-7.32 7.32h-8.708V23.052h8.708c4.036 0 7.32 3.282 7.32 7.316v11.26Z"/></defs><use href="#openmojiRegionalIndicatorD0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorD0"/></g>`),
		g.Group(children),
	)
}