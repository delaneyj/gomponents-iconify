package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorM0" d="M48.23 21.04a1 1 0 0 0-1.124.525L36 43.776L24.895 21.565a1 1 0 0 0-1.894.447v28a1 1 0 1 0 1.999 0V26.248l10.106 20.211c.34.678 1.449.678 1.789 0L47 26.248v23.764a1 1 0 1 0 2 0v-28a1 1 0 0 0-.77-.973Z"/></defs><use href="#openmojiRegionalIndicatorM0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorM0"/></g>`),
		g.Group(children),
	)
}