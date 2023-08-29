package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareFEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M18.333 2c1.96 0 3.56 1.537 3.662 3.472l.005.195v12.666c0 1.96-1.537 3.56-3.472 3.662l-.195.005H5.667a3.667 3.667 0 0 1-3.662-3.472L2 18.333V5.667c0-1.96 1.537-3.56 3.472-3.662L5.667 2h12.666zM15 8h-1l-.15.005a2 2 0 0 0-1.844 1.838L12 10v1l.005.15c.018.236.077.46.17.667l.075.152l.018.03l-.018.032c-.133.24-.218.509-.243.795L12 13v1l.005.15a2 2 0 0 0 1.838 1.844L14 16h1l.15-.005a2 2 0 0 0 1.844-1.838L17 14v-1l-.005-.15a1.988 1.988 0 0 0-.17-.667l-.075-.152l-.019-.032l.02-.03a2.01 2.01 0 0 0 .242-.795L17 11v-1l-.005-.15a2 2 0 0 0-1.838-1.844L15 8zm-5 0H8l-.117.007a1 1 0 0 0-.876.876L7 9v6l.007.117a1 1 0 0 0 .876.876L8 16l.117-.007a1 1 0 0 0 .876-.876L9 15v-2h1l.117-.007a1 1 0 0 0 0-1.986L10 11H9v-1h1l.117-.007a1 1 0 0 0 0-1.986L10 8zm5 5v1h-1v-1h1zm0-3v1h-1v-1h1z"/></g>`),
		g.Group(children),
	)
}