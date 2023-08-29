package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M350 249q30 17 47 47t17 63q-29 17-63 17.5T286 359q-9-5-17-11q2 10 2 20q0 35-17.5 64.5T207 479q-29-17-46.5-46.5T143 368q0-10 2-20q-9 7-17 11q-31 17-65 17T0 359q0-34 17-63.5T64 248q8-4 18-8q-10-4-18-9q-30-17-47-47T0 121q29-17 63-17.5t65 17.5q8 4 17 11q-2-10-2-20q0-35 17.5-64.5T207 1q29 17 46.5 46.5T271 112q0 10-2 20q9-7 17-11q31-18 65-17.5t63 17.5q0 33-17 63t-47 47q-8 5-18 9q10 4 18 9zm-143 76q35 0 60-25t25-60t-25-60t-60-25t-60 25t-25 60t25 60t60 25z"/>`),
		g.Group(children),
	)
}