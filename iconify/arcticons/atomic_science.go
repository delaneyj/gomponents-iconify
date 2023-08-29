package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtomicScience(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.5 15.5h-9a2 2 0 0 1-2-2v-9h-18a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h27a2 2 0 0 0 2-2v-26Zm-11-11l11 11"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="24" cy="28.247" rx="4.822" ry="11.5"/><ellipse cx="24" cy="28.247" rx="11.5" ry="4.822" transform="rotate(-30 24 28.247)"/><ellipse cx="24" cy="28.247" rx="4.822" ry="11.5" transform="rotate(-60 24 28.247)"/></g>`),
		g.Group(children),
	)
}