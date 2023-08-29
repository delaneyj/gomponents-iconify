package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Electricity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.07 26.285c.962 1.176 3.41 5.533 3.93 7.715h14c.52-2.18 2.965-6.537 3.927-7.712a13.568 13.568 0 0 0 2.982-7.015a13.47 13.47 0 0 0-1.289-7.495a13.813 13.813 0 0 0-5.164-5.671A14.215 14.215 0 0 0 24.002 4c-2.637 0-5.221.73-7.454 2.105a13.814 13.814 0 0 0-5.166 5.67a13.47 13.47 0 0 0-1.292 7.493c.299 2.567 1.332 5 2.98 7.017ZM25 12l-6 9h4v6l6-9h-4v-6Z" clip-rule="evenodd"/><path d="M17 37a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H18a1 1 0 0 1-1-1Zm14 3H17v2a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-2Z"/></g>`),
		g.Group(children),
	)
}