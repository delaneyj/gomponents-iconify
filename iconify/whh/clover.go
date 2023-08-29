package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clover(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M791 485q149 54 202 105q14 14 22.5 38t8.5 55t-17 67.5t-51 70.5q-45 44-100.5 62.5t-96 12.5t-53.5-28q-25-42-72-114t-84-131q10 92 18 203t8 135q0 27-18.5 45.5t-45 18.5t-45.5-18.5t-19-45.5q0-25 8-140.5T475 613q-38 58-86 128.5T317 852q-11 16-36.5 28T222 896t-73.5-9.5t-75-45.5t-53-73t-20-77t6-65T28 584q31-29 99.5-60.5T279 466q-43-15-96.5-30t-91-27.5T22 377Q6 367 2.5 336T8 266t33-84t58-80.5t77.5-60t81-34t68-6.5T368 18q83 102 142 291q46-132 136-286q9-17 44-20t79 6.5t93.5 34T947 102q34 34 53.5 77t22 80t-4.5 66.5t-21 43.5q-57 56-206 116z"/>`),
		g.Group(children),
	)
}