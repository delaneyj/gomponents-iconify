package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductRelease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 1537V936L256 616v880l544 273l-31 127l-641-320V472L960 57l832 415v270q-70 11-128 45V616l-640 320v473l-128 128zM754 302l584 334l247-124l-625-313l-206 103zm206 523l240-120l-584-334l-281 141l625 313zm888 71q42 0 78 15t64 41t42 63t16 79q0 39-15 76t-43 65l-717 717l-377 94l94-377l717-716q29-29 65-43t76-14zm51 249q21-21 21-51q0-31-20-50t-52-20q-14 0-27 4t-23 15l-692 692l-34 135l135-34l692-691z"/>`),
		g.Group(children),
	)
}