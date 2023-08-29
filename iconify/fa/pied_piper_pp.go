package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PiedPiperPp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1046 892q0 64-38 109t-91 45q-43 0-70-15V754q28-17 70-17q53 0 91 45.5t38 109.5zM703 464q0 64-38 109.5T574 619q-43 0-70-15V327q28-17 70-17q53 0 91 45t38 109zm562 431q0-134-88-229t-213-95q-20 0-39 3q-23 78-78 136q-87 95-211 101v636l211-41v-206q51 19 117 19q125 0 213-95t88-229zM922 468q0-134-88.5-229T620 144q-74 0-141 36H293v840l211-41V773q55 19 116 19q125 0 213.5-95T922 468zm614-180v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288z"/>`),
		g.Group(children),
	)
}