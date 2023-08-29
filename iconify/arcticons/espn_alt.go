package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EspnAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 29.823h5.823m2.053-11.646H6.553l-.513 2.911M5.527 24h3.796m-3.796 0L4.5 29.823M24.312 26h3.813c2.155 0 4.211-1.752 4.592-3.912s-1.057-3.911-3.213-3.911h-3.812l-.514 2.911m-1.54 8.735L24.665 24m17.249 5.823L42.941 24m.514-2.912h0c.283-1.608-.79-2.91-2.399-2.91h-4.804l-.513 2.91m-1.54 8.735L35.226 24m-20.763-2.912c.283-1.607 1.814-2.91 3.418-2.91h1.724c1.246 0 2.08.345 2.631 1.275m-8.858 9.094c.55.93 1.385 1.276 2.63 1.276h1.725c1.604 0 3.135-1.303 3.419-2.911h0C21.435 25.304 20.364 24 18.76 24h-1.905"/>`),
		g.Group(children),
	)
}