package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bonbon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="5.406" height="7.163" x="21.308" y="14.649" rx="2.703" ry="2.703"/><path d="M29.343 21.812v-4.46a2.703 2.703 0 0 1 2.703-2.703h0a2.703 2.703 0 0 1 2.703 2.703v4.46m-21.498-4.46a2.703 2.703 0 0 1 2.703-2.703h0a2.703 2.703 0 0 1 2.703 2.703v1.757a2.703 2.703 0 0 1-2.703 2.703h0a2.703 2.703 0 0 1-2.703-2.703m0 0V11"/><rect width="5.406" height="7.163" x="21.308" y="27.837" rx="2.703" ry="2.703"/><path d="M29.343 35v-4.46a2.703 2.703 0 0 1 2.703-2.703h0a2.703 2.703 0 0 1 2.703 2.703V35m-21.498-4.46a2.703 2.703 0 0 1 2.703-2.703h0a2.703 2.703 0 0 1 2.703 2.703v1.757A2.703 2.703 0 0 1 15.954 35h0a2.703 2.703 0 0 1-2.703-2.703m0 0v-8.109"/></g>`),
		g.Group(children),
	)
}