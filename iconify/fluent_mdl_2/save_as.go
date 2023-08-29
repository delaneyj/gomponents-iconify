package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveAs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1848 896q42 0 78 15t64 41t42 63t16 79q0 39-15 76t-43 65l-717 717l-377 94l94-377l717-716q29-29 65-43t76-14zm51 249q21-21 21-51q0-31-20-50t-52-20q-14 0-27 4t-23 15l-692 692l-34 135l135-34l692-691zM768 1536h128l-128 128H293l-165-165V256q0-27 10-50t27-40t41-28t50-10h1280q27 0 50 10t40 27t28 41t10 50v512l-128 128V256h-128v640H384V256H256v1189l91 91h37v-512h896v128l-128 128v-128H512v384h128v-256h128v256zM512 768h768V256H512v512z"/>`),
		g.Group(children),
	)
}