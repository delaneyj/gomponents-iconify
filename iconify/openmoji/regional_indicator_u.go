package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorU(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorU0" d="M46 21a1 1 0 0 0-1 1v18c0 4.962-4.037 9-9 9s-9-4.038-9-9V22a1 1 0 1 0-2 0v18c0 6.065 4.935 11 11 11s11-4.935 11-11V22a1 1 0 0 0-1-1Z"/></defs><use href="#openmojiRegionalIndicatorU0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorU0"/></g>`),
		g.Group(children),
	)
}