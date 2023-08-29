package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.958 9.221V1.783S13.14.052 10.896.052H6.084c-2.244 0-2.062 1.731-2.062 1.731v7.438s-.146 1.748 1.062 1.748h1.959v.948h1.001V14H7.043v1h1.001v.917H9V15h.938v-1H9v-2.083h1v-.948h1.896c1.354 0 1.062-1.748 1.062-1.748zm-2.308.946H10V10H7.042v.167H6.39c-1.546 0-1.478-1.356-1.478-1.356V4.954h2.057v-.901H4.912V2.954h2.041v-.901H4.965c.104-.424.423-1.136 1.424-1.136h4.261c1.35 0 1.479 1.445 1.479 1.445v6.449c.001 0 .121 1.356-1.479 1.356z"/><path d="M6.023 8.137c0 .696.703.807.703.807h3.541s.704-.026.704-.807v-3.11H6.023v3.11z"/></g>`),
		g.Group(children),
	)
}