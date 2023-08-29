package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorO0" d="M36 21.012c-6.065 0-11 4.935-11 11v8c0 6.066 4.935 11 11 11s11-4.934 11-11v-8c0-6.065-4.935-11-11-11Zm9 19c0 4.963-4.037 9-9 9s-9-4.037-9-9v-8c0-4.962 4.037-9 9-9s9 4.038 9 9v8Z"/></defs><use href="#openmojiRegionalIndicatorO0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorO0"/></g>`),
		g.Group(children),
	)
}