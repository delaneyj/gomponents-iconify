package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldTimes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="m1000 20l1000 300q0 136-2.5 228t-13 217.5t-30 216t-55 205T1813 1391t-126.5 195t-172.5 196t-226.5 188.5T1000 2160q-158-92-287.5-189.5T486 1782t-172.5-196T187 1391t-86.5-204.5t-55-205t-30-216T2.5 548T0 320zm497 1237q0-34-25-59l-236-236l236-235q25-25 25-60q0-36-25-59l-117-118q-23-24-59-24t-60 24l-236 237l-236-237q-24-24-59-24q-37 0-59 24L528 608q-26 24-26 59q0 34 26 60l236 235l-236 236q-26 26-26 59q0 36 26 60l118 118q23 25 59 25q34 0 59-25l236-237l236 237q25 25 60 25q36 0 59-25l117-118q25-23 25-60z"/>`),
		g.Group(children),
	)
}