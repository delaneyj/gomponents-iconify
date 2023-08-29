package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2zm1 5h-2l-.15.005a2 2 0 0 0-1.844 1.838L9 9v2l.005.15c.018.236.077.46.17.667l.075.152l.018.03l-.018.032c-.133.24-.218.509-.243.795L9 13v2l.005.15a2 2 0 0 0 1.838 1.844L11 17h2l.15-.005a2 2 0 0 0 1.844-1.838L15 15v-2l-.005-.15a1.988 1.988 0 0 0-.17-.667l-.075-.152l-.019-.032l.02-.03a2.01 2.01 0 0 0 .242-.795L15 11V9l-.005-.15a2 2 0 0 0-1.838-1.844L13 7zm0 6v2h-2v-2h2zm0-4v2h-2V9h2z"/></g>`),
		g.Group(children),
	)
}