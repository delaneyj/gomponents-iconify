package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HintText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v384h-128V384H128v1152h896v128H0V256h2048zm-512 384q106 0 199 40t163 109t110 163t40 200q0 106-38 195t-111 166l-32 34q-12 13-25 29t-23 36t-17 38t-7 38v168q0 41-15 76t-42 62t-62 41t-76 16h-128q-41 0-76-15t-62-42t-41-61t-16-77v-168q-1-25-11-49t-25-46t-33-42t-35-38q-72-76-110-165t-39-196q0-106 40-199t109-163t163-110t200-40zm125 1155h-250v61q0 25 18 43t43 18h128q25 0 43-18t18-43v-61zm9-134q10-57 32-98t50-76t57-66t54-69t41-85t16-115q0-80-30-150t-82-122t-122-82t-150-30q-80 0-150 30t-122 82t-82 122t-30 150q0 76 29 147t83 128q27 27 49 52t40 53t30 58t19 71h268z"/>`),
		g.Group(children),
	)
}