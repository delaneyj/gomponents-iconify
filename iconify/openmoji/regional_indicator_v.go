package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorV0" d="M43.274 21.038a1.003 1.003 0 0 0-1.236.687L35 46.36l-7.038-24.635a1 1 0 0 0-1.924.55l8 28a1 1 0 0 0 1.924 0l8-28a1 1 0 0 0-.688-1.236Z"/></defs><use href="#openmojiRegionalIndicatorV0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorV0"/></g>`),
		g.Group(children),
	)
}