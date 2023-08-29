package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellPhone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 0q27 0 50 10t40 27t28 41t10 50v1792q0 27-10 50t-27 40t-41 28t-50 10H512q-27 0-50-10t-40-27t-28-41t-10-50V128q0-27 10-50t27-40t41-28t50-10h1024zm0 128H512v1792h1024V128zM896 1664h256v128H896v-128z"/>`),
		g.Group(children),
	)
}