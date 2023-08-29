package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GbEng(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<mask id="flagpackGbEng0" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse"><path fill="#fff" d="M0 0h32v24H0z"/></mask><g fill="none" fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackGbEng0)"><path fill="#F7FCFF" d="M0 0v24h32V0H0z"/><path fill="#F50302" d="M18 0h-4v10H0v4h14v10h4V14h14v-4H18V0z"/></g>`),
		g.Group(children),
	)
}