package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorE(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorE0" d="M45 23a1 1 0 0 0 0-2H29a1 1 0 0 0-1 1v28a1 1 0 0 0 1 1h16a1 1 0 0 0 0-2H30V37h11a1 1 0 0 0 0-2H30V23h15Z"/></defs><use href="#openmojiRegionalIndicatorE0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorE0"/></g>`),
		g.Group(children),
	)
}