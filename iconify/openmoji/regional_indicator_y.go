package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorY(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorY0" d="M44.447 21.105a1 1 0 0 0-1.341.448L36 35.764l-7.105-14.211a1 1 0 1 0-1.79.894L35 38.236V50a1 1 0 1 0 2 0V38.236l7.895-15.789a1 1 0 0 0-.447-1.342Z"/></defs><use href="#openmojiRegionalIndicatorY0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorY0"/></g>`),
		g.Group(children),
	)
}