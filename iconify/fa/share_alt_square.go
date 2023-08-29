package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareAltSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1280 1067q0-88-62.5-151T1067 853q-84 0-145 58L681 791q2-16 2-23t-2-23l241-120q61 58 145 58q88 0 150.5-63t62.5-151t-62.5-150.5T1067 256t-151 62.5T853 469q0 7 2 23L614 612q-62-57-145-57q-88 0-150.5 62.5T256 768t62.5 150.5T469 981q83 0 145-57l241 120q-2 16-2 23q0 88 63 150.5t151 62.5t150.5-62.5T1280 1067zm256-779v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288z"/>`),
		g.Group(children),
	)
}