package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorP(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorP0" d="M40.85 21.012H30.174a1 1 0 0 0-1 1v28a1 1 0 1 0 2 0V36.96h9.676c4.398 0 7.976-3.576 7.976-7.972s-3.578-7.976-7.976-7.976Zm0 13.948h-9.676V23.012h9.676a5.983 5.983 0 0 1 5.976 5.976a5.981 5.981 0 0 1-5.976 5.972Z"/></defs><use href="#openmojiRegionalIndicatorP0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorP0"/></g>`),
		g.Group(children),
	)
}