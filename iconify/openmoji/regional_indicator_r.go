package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorR0" d="M47.826 28.988c0-4.398-3.578-7.976-7.975-7.976H29.174a1 1 0 0 0-1 1v28a1 1 0 1 0 2 0V36.96h9.261l5.81 13.446a1 1 0 1 0 1.835-.793l-5.545-12.835c3.59-.776 6.291-3.972 6.291-7.79Zm-7.975 5.972h-9.677V23.012h9.677a5.983 5.983 0 0 1 5.975 5.976a5.98 5.98 0 0 1-5.975 5.973Z"/></defs><use href="#openmojiRegionalIndicatorR0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorR0"/></g>`),
		g.Group(children),
	)
}