package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorH0" d="M44 20.508a1 1 0 0 0-1 1v12.96H29v-12.96a1 1 0 1 0-2 0v28a1 1 0 1 0 2 0v-13.04h14v13.04a1 1 0 1 0 2 0v-28a1 1 0 0 0-1-1Z"/></defs><use href="#openmojiRegionalIndicatorH0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorH0"/></g>`),
		g.Group(children),
	)
}