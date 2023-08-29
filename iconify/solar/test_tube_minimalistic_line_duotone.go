package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTubeMinimalisticLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M21 9.849L14.181 3m.683.685l5.375 5.399L13 16.354l-3.512 3.528a3.789 3.789 0 0 1-5.375 0a3.83 3.83 0 0 1 0-5.4l10.75-10.797Z"/><path d="m6.8 11.783l1.275.142a2.205 2.205 0 0 1 1.944 1.953a2.209 2.209 0 0 0 1.32 1.787l1.661.69" opacity=".5"/></g>`),
		g.Group(children),
	)
}