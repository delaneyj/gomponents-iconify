package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchwabischGmundForumGoldUndSilber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9b9b9a" d="M20 40h35v11H20z"/><path fill="#fcea2b" d="m46 18l13 4l-3 20l-5 5l-3-5l-2-24Z"/><path fill="#fcea2b" d="m15 17l4 14l-1 16l30-5l-2-24Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 17l4 14l-1 16l30-5l-2-24Zm4 14l27-13m2 24L19 31m27-13l10 24m-2 9H20v-4"/><path d="m46 18l13 4l-3 20l-5 5l-3-5l-2-24Zm0 0l10 24"/></g>`),
		g.Group(children),
	)
}