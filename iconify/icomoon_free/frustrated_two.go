package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrustratedTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 10.25v1.5c0 .136.114.25.25.25H6v-2H4.25a.253.253 0 0 0-.25.25zM7 10h2v2H7v-2zm4.75 0H10v2h1.75c.136 0 .25-.114.25-.25v-1.5a.253.253 0 0 0-.25-.25z"/><path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm1.002 5.946c.032-.741.706-1.234 1.275-1.518c.543-.271 1.08-.407 1.102-.413a.5.5 0 1 1 .242.97c-.275.07-.602.189-.89.334A.998.998 0 0 1 9.999 7a1 1 0 0 1-1-1l.002-.054zM4.015 4.379a.5.5 0 0 1 .606-.364c.023.006.559.141 1.102.413c.568.284 1.243.776 1.275 1.518L7 6a1 1 0 1 1-1.732-.681a4.638 4.638 0 0 0-.89-.334a.5.5 0 0 1-.364-.606zM13 11.75c0 .689-.561 1.25-1.25 1.25h-7.5C3.561 13 3 12.439 3 11.75v-1.5C3 9.561 3.561 9 4.25 9h7.5c.689 0 1.25.561 1.25 1.25v1.5z"/>`),
		g.Group(children),
	)
}