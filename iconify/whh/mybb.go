package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mybb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 416q0 30 5 57t13 45.5t17 32t16.5 19.5t12.5 6q-74 0-138.5-19T776 505q-37 7-72.5 7t-71.5-7l-2 4l-1 3h11q87 0 160.5 30T917 623.5T960 736t-43 112.5T800.5 930T640 960q-36 0-73-6q-45 33-109.5 51.5T320 1024q9 0 23-16t27-54t14-84q-64-60-64-134q0-14 4-32h-4q-35 0-72-7q-45 33-109.5 52T0 768q5 0 12.5-6T29 742.5t17-32T59 665t5-57q0-1 .5-3t.5-3Q0 533 0 448q0-69 43-128t116.5-93.5T320 192q36 0 72 7q25-86 112-142.5T704 0q87 0 160.5 34.5T981 128t43 128q0 85-65 154q0 1 .5 3t.5 3z"/>`),
		g.Group(children),
	)
}