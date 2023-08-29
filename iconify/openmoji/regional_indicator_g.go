package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorG(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorG0" d="M46 34.448h-8a1 1 0 0 0 0 2h7v3.056c0 4.963-4.038 9-9 9s-9-4.037-9-9v-8c0-4.962 4.038-9 9-9c2.209 0 4.33.804 5.972 2.264a1 1 0 1 0 1.329-1.495A10.975 10.975 0 0 0 36 20.504c-6.065 0-11 4.935-11 11v8c0 6.066 4.935 11 11 11s11-4.934 11-11v-4.056a1 1 0 0 0-1-1Z"/></defs><use href="#openmojiRegionalIndicatorG0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorG0"/></g>`),
		g.Group(children),
	)
}