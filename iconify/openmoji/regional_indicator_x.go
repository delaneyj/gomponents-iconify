package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorX0" d="M44.495 21.122a.998.998 0 0 0-1.364.372l-7.137 12.49l-7.125-12.468a1 1 0 1 0-1.736.992L34.842 36l-7.71 13.492a1 1 0 1 0 1.737.992l7.125-12.469l7.137 12.49a.999.999 0 0 0 1.364.373a1 1 0 0 0 .372-1.364L37.145 36l7.722-13.514a1 1 0 0 0-.372-1.364Z"/></defs><use href="#openmojiRegionalIndicatorX0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorX0"/></g>`),
		g.Group(children),
	)
}