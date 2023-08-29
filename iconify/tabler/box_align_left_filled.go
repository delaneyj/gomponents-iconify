package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxAlignLeftFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M10.002 3.003h-5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h5a1 1 0 0 0 1-1v-16a1 1 0 0 0-1-1zm5 16a1 1 0 0 1 .117 1.993l-.127.007a1 1 0 0 1-.117-1.993l.127-.007zm5.001 0a1 1 0 0 1 .117 1.993l-.128.007a1 1 0 0 1-.117-1.993l.128-.007zm0-5.001a1 1 0 0 1 .117 1.993l-.128.007a1 1 0 0 1-.117-1.993l.128-.007zm0-6a1 1 0 0 1 .117 1.993l-.128.007a1 1 0 0 1-.117-1.993l.128-.007zm0-5a1 1 0 0 1 .117 1.993l-.128.007a1 1 0 0 1-.117-1.993l.128-.007zm-5.001 0a1 1 0 0 1 .117 1.993l-.127.007a1 1 0 0 1-.117-1.993l.127-.007z"/></g>`),
		g.Group(children),
	)
}