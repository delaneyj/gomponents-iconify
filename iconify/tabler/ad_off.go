package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 5h10a2 2 0 0 1 2 2v10m-2 2H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2"/><path d="M7 15v-4a2 2 0 0 1 2-2m2 2v4m-4-2h4m6-4v4m-.885-.869c.33.149.595.412.747.74M3 3l18 18"/></g>`),
		g.Group(children),
	)
}