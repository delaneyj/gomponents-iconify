package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 6h-.975A2.025 2.025 0 0 0 3 8.025v2.834c0 .537.213 1.052.593 1.432l6.116 6.116a2.025 2.025 0 0 0 2.864 0l2.834-2.834c.028-.028.055-.056.08-.085m2.086 2.919l.418-.418m2-2l.419-.419a2.025 2.025 0 0 0 0-2.864L13.293 5.59M6 9h-.01M3 3l18 18"/>`),
		g.Group(children),
	)
}