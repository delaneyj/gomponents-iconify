package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenEnrollment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1848 896q42 0 78 15t64 42t42 63t16 78q0 39-15 76t-43 65l-717 719l-377 94l94-377l717-718q28-28 65-42t76-15zm51 249q21-21 21-51q0-31-20-50t-52-20q-14 0-27 4t-23 15l-692 694l-34 135l135-34l692-693zM640 896H512V768h128v128zm896 0H768V768h768v128zM512 1152h128v128H512v-128zm128-640H512V384h128v128zm896 0H768V384h768v128zM384 1664h443l-32 128H256V0h1536v743q-67 10-128 44V128H384v1536zm384-512h514l-128 128H768v-128z"/>`),
		g.Group(children),
	)
}