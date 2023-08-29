package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopSecure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M884 1536q15 67 42 128H128q-27 0-50-10t-40-27t-28-41t-10-50q0-16 3-34t8-38t14-35t21-29l210-211V256h1536v611q-30-17-62-28t-66-19V384H384v768h480v128H347l-206 205q-3 3-5 10t-4 14t-3 15t-1 12H64h820zm652-571q59 0 108 14t100 43q42 25 96 31t108 7h51q25 0 49-1v325q0 105-38 198t-101 173t-145 144t-169 110l-59 31l-60-31q-87-45-168-109t-145-145t-101-173t-38-198v-325h50q26 0 54 1q53 0 106-6t96-33q50-29 99-42t107-14zm384 222q-66 0-127-12t-118-45q-34-20-66-28t-73-9q-38 0-71 8t-66 28q-58 33-119 45t-128 13v197q0 86 33 162t88 142t123 118t140 90q70-37 139-89t123-118t88-142t34-163v-197z"/>`),
		g.Group(children),
	)
}