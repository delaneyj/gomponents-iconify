package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxAlignTopFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M19 3.005H5a2 2 0 0 0-2 2v5a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-5a2 2 0 0 0-2-2zM4 13.995a1 1 0 0 1 .993.883l.007.127a1 1 0 0 1-1.993.117L3 14.995a1 1 0 0 1 1-1zm0 5a1 1 0 0 1 .993.883l.007.127a1 1 0 0 1-1.993.117L3 19.995a1 1 0 0 1 1-1zm5 0a1 1 0 0 1 .993.883l.007.127a1 1 0 0 1-1.993.117L8 19.995a1 1 0 0 1 1-1zm6 0a1 1 0 0 1 .993.883l.007.127a1 1 0 0 1-1.993.117L14 19.995a1 1 0 0 1 1-1zm5 0a1 1 0 0 1 .993.883l.007.127a1 1 0 0 1-1.993.117L19 19.995a1 1 0 0 1 1-1zm0-5a1 1 0 0 1 .993.883l.007.127a1 1 0 0 1-1.993.117L19 14.995a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}