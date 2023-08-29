package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorW(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorW0" d="M52.275 21.05a1.002 1.002 0 0 0-1.236.687l-7.037 24.635l-7.042-24.636a1 1 0 0 0-1.924 0L28 46.37l-7.038-24.633a1 1 0 0 0-1.924.55l8 27.998a1 1 0 0 0 1.924 0l7.037-24.635l7.042 24.636a1 1 0 0 0 1.924 0l7.999-28a1 1 0 0 0-.688-1.236Z"/></defs><use href="#openmojiRegionalIndicatorW0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorW0"/></g>`),
		g.Group(children),
	)
}