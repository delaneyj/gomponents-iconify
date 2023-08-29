package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timeline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1110q63 77 95 169t33 193q0 119-45 224t-124 183t-183 123t-224 46q-119 0-224-45t-183-124t-123-183t-46-224q0-100 33-192H0V256h256v128H128v256h768V384H768V256h384v128h-128v256h768V384h-128V256h256v854zm-927 42q40-60 93-107t114-81t131-50t141-18q86 0 167 24t153 73V768H128v384h865zm479 768q93 0 174-35t142-96t96-142t36-175q0-93-35-174t-96-142t-142-96t-175-36q-93 0-174 35t-142 96t-96 142t-36 175q0 93 35 174t96 142t142 96t175 36zm64-512h192v128h-320v-384h128v256z"/>`),
		g.Group(children),
	)
}