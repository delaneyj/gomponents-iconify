package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChangeEntitlements(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 1920h899l128 128H0V0h1115l549 549v219h-128V640h-512V128H128v1792zM1152 219v293h293l-293-293zM256 512h639v127H256V512zm1151 256v127H256V768h1151zM256 1024h767v127H256v-127zm768 256v127H256v-127h768zm-768 383v-127h768l-128 127H256zm1213-162l-163 162h614v129h-614l163 163l-90 90l-318-317l318-317l90 90zm297-349l-163-163l90-90l318 317l-318 317l-90-90l163-163h-614v-128h614z"/>`),
		g.Group(children),
	)
}