package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefreshSyncAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 192q0-59 36-105t92-60v44q-38 14-61.5 47T43 192q0 53 37 90l48-47v128H0l50-51Q0 262 0 192zm171 107v-43h42v43h-42zM384 21l-50 51q50 50 50 120q0 59-36 105t-92 60v-44q38-14 61.5-47t23.5-74q0-53-37-90l-48 47V21h128zM171 213V85h42v128h-42z"/>`),
		g.Group(children),
	)
}