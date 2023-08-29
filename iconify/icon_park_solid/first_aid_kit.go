package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstAidKit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFirstAidKit0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M8 20v18a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V20"/><path fill="#fff" d="M5 14h38v6H5v-6Zm26-6H17v6h14V8Z"/><path stroke-linecap="round" d="M20 30h8m-4-4v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFirstAidKit0)"/>`),
		g.Group(children),
	)
}