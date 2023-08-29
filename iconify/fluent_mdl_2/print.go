package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 768q26 0 49 10t41 27t28 41t10 50v896h-512v256H512v-256H0V896q0-26 10-49t27-41t41-28t50-10h384V0h1024v768h384zm-1280 0h768V128H640v640zm768 640H640v512h768v-512zm512-512H128v768h384v-384h1024v384h384V896zM320 1024q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45q0-26 19-45t45-19z"/>`),
		g.Group(children),
	)
}