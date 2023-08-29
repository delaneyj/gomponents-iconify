package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EatDrink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M960 0q26 0 45 19t19 45v704q0 53-20 99t-55 81t-82 55t-99 21v960q0 26-19 45t-45 19q-26 0-45-19t-19-45v-960q-53 0-99-20t-81-55t-55-81t-21-100V64q0-26 19-45t45-19q26 0 45 19t19 45v704q0 27 10 50t27 40t41 28t50 10V64q0-26 19-45t45-19q26 0 45 19t19 45v832q27 0 50-10t40-27t28-41t10-50V64q0-26 19-45t45-19zm704 0v1984q0 26-19 45t-45 19q-26 0-45-19t-19-45v-576q-37 0-80 1t-85-1t-82-12t-70-31t-49-57t-18-92V448q0-93 35-174t96-142t142-96t175-36h64zm-128 134q-56 11-102 40t-81 72t-54 93t-19 109v768q0 26 19 45t45 19h192V134z"/>`),
		g.Group(children),
	)
}