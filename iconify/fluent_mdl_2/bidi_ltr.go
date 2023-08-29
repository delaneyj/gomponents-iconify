package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BidiLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m128 486l666 666l-666 666V486zm128 1024l358-358l-358-358v716zM1792 256v1536h128v128h-640v-128h128v-768h-192q-93 0-174-35t-143-96t-96-142t-35-175q0-93 35-174t96-143t142-96t175-35h704v128h-128zm-384 640V256h-192q-66 0-124 25t-102 69t-69 102t-25 124q0 66 25 124t68 102t102 69t125 25h192zm256 896V256h-128v1536h128z"/>`),
		g.Group(children),
	)
}