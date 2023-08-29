package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxAlignBottomRightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M19 12h-5a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2v-5a2 2 0 0 0-2-2zm1-4a1 1 0 0 1 .993.883L21 9.01a1 1 0 0 1-1.993.117L19 9a1 1 0 0 1 1-1zm0-5a1 1 0 0 1 .993.883L21 4.01a1 1 0 0 1-1.993.117L19 4a1 1 0 0 1 1-1zm-5 0a1 1 0 0 1 .993.883L16 4.01a1 1 0 0 1-1.993.117L14 4a1 1 0 0 1 1-1zM9 3a1 1 0 0 1 .993.883L10 4.01a1 1 0 0 1-1.993.117L8 4a1 1 0 0 1 1-1zm0 16a1 1 0 0 1 .993.883l.007.127a1 1 0 0 1-1.993.117L8 20a1 1 0 0 1 1-1zM4 3a1 1 0 0 1 .993.883L5 4.01a1 1 0 0 1-1.993.117L3 4a1 1 0 0 1 1-1zm0 5a1 1 0 0 1 .993.883L5 9.01a1 1 0 0 1-1.993.117L3 9a1 1 0 0 1 1-1zm0 6a1 1 0 0 1 .993.883L5 15.01a1 1 0 0 1-1.993.117L3 15a1 1 0 0 1 1-1zm0 5a1 1 0 0 1 .993.883L5 20.01a1 1 0 0 1-1.993.117L3 20a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}