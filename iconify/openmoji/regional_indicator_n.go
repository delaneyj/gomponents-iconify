package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorN(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorN0" d="M46 21.012a1 1 0 0 0-1 1v24.88l-18.186-25.46a1 1 0 0 0-1.814.58v28a1 1 0 1 0 2 0v-24.88l18.187 25.461a1 1 0 0 0 1.813-.58v-28a1 1 0 0 0-1-1Z"/></defs><use href="#openmojiRegionalIndicatorN0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorN0"/></g>`),
		g.Group(children),
	)
}