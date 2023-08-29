package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Edge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M5 795h1q16-126 58.5-241.5t115-217t167.5-176T570.5 43T847 0q231 0 414 105.5T1555 409q104 187 104 442v188H534q1 111 53.5 192.5T724 1354t189.5 57t213 3t208-46.5T1508 1283v377q-92 55-229.5 92T966 1790t-316-53q-189-73-311.5-249T214 1116q-3-242 111-412t325-268q-48 60-78 125.5T526 721h635q8-77-8-140t-47-101.5t-70.5-66.5t-80.5-41t-75-20.5t-56-8.5l-22-1q-135 5-259.5 44.5T319 491T143 631.5T5 795z"/>`),
		g.Group(children),
	)
}