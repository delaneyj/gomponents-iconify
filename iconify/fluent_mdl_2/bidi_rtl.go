package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BidiRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1152 128v128h-128v1536h128v128H512v-128h128v-768H448q-93 0-174-35t-143-96t-96-142T0 576q0-93 35-174t96-143t142-96t175-35h704zM640 896V256H448q-66 0-124 25t-102 69t-69 102t-25 124q0 66 25 124t68 102t102 69t125 25h192zm256 896V256H768v1536h128zM1920 486v1332l-666-666l666-666zm-128 308l-358 358l358 358V794z"/>`),
		g.Group(children),
	)
}