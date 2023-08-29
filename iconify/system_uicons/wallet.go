package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" transform="translate(3 4)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 2.5h12a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2zm1-2h9a1 1 0 0 1 1 1v1H.5v-1a1 1 0 0 1 1-1z"/><circle cx="11.5" cy="7.5" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}