package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReligionHexagramStarJewJewishJudaismHexagramCultureReligionDavid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 3.5h13L7 13.5L.5 3.5z"/><path d="M.5 10.5h13L7 .5l-6.5 10z"/></g>`),
		g.Group(children),
	)
}