package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M93 219q0 46 28 74q32 32 81 32h7q57 0 89-32q30-30 30-76q0-33-16.5-64T273 100.5t-47.5-39t-40-25T166 27L117 5l19 49q7 25 4.5 35.5T125 118q-2 4-10 19.5t-11 24t-7 25t-4 32.5zm69-75q21-31 23-55q100 60 100 130q0 31-17 44q-8 9-23 13q7-19 5.5-39.5t-4.5-32t-5-15.5q-9-20-28-13q-8 3-12 11.5t-1 16.5q17 53-9 77q-18 0-36.5-15T136 219q0-7 .5-13.5t1.5-11t3-10.5t3-8.5t4.5-8.5t4-7.5t4.5-8t5-7.5zm245 160l-196 53l-197-53l-12 43l125 32L2 411l12 42l197-53l196 53l13-42l-126-32l126-32z"/>`),
		g.Group(children),
	)
}