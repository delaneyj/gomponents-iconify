package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorF0" d="M46 21H30a1 1 0 0 0-1 1v28a1 1 0 0 0 2 0V37h11a1 1 0 0 0 0-2H31V23h15a1 1 0 0 0 0-2Z"/></defs><use href="#openmojiRegionalIndicatorF0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorF0"/></g>`),
		g.Group(children),
	)
}