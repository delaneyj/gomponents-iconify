package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M896 128q-204 0-381.5 69.5T232.5 385T128 640q0 112 71.5 213.5T401 1029l87 50l-27 96q-24 91-70 172q152-63 275-171l43-38l57 6q69 8 130 8q204 0 381.5-69.5t282-187.5T1664 640t-104.5-255t-282-187.5T896 128zm896 512q0 174-120 321.5t-326 233t-450 85.5q-70 0-145-8q-198 175-460 242q-49 14-114 22h-5q-15 0-27-10.5t-16-27.5v-1q-3-4-.5-12t2-10t4.5-9.5l6-9l7-8.5l8-9q7-8 31-34.5t34.5-38t31-39.5t32.5-51t27-59t26-76q-157-89-247.5-220T0 640q0-174 120-321.5t326-233T896 0t450 85.5t326 233T1792 640z"/>`),
		g.Group(children),
	)
}