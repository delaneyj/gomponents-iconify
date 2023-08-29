package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blogger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M397 165q30 0 30 31v103q0 53-39.5 91.5T295 429H124q-48 0-86-38.5T0 300V138q0-57 39-96t96-39h90q44 0 84.5 39.5T350 128v11q0 11 7.5 18.5T378 165h19zm-262-51q-10 0-17.5 7.5t-7.5 18t7.5 18T135 165h78q10 0 17-8t7-18t-7-17.5t-17-7.5h-78zm154 204q10 0 17.5-6.5T314 295t-7.5-17t-17.5-7H135q-10 0-17.5 7t-7.5 17t7.5 16.5T135 318h154z"/>`),
		g.Group(children),
	)
}