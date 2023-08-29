package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchwbischGmndForumGoldUndSilber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9b9b9a" d="M20 40h35v11H20z"/><path fill="#fcea2b" d="M46 18l13 4l-3 20l-5 5l-3-5l-2-24z"/><path fill="#fcea2b" d="M15 17l4 14l-1 16l30-5l-2-24z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 17l4 14l-1 16l30-5l-2-24z"/><path d="M19 31l27-13"/><path d="M48 42L19 31"/><path d="M46 18l10 24"/><path d="M54 51H20v-4"/><path d="M46 18l13 4l-3 20l-5 5l-3-5l-2-24z"/><path d="M46 18l10 24"/></g>`),
		g.Group(children),
	)
}