package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscGolf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5h14M6 5c.32 6.744 2.74 9.246 6 10m6-10c-.32 6.744-2.74 9.246-6 10M10 5c0 4.915.552 7.082 2 10m2-10c0 4.915-.552 7.082-2 10m0 0v6m0-18v2M7 16c.64.64 1.509 1 2.414 1h5.172c.905 0 1.774-.36 2.414-1m-6 5h2"/>`),
		g.Group(children),
	)
}