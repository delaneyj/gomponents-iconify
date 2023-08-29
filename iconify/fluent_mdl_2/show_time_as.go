package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShowTimeAs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 0h2048v2048H0V0zm128 768v165l165-165H128zm0 347v165h165l512-512H475l-347 347zm1792 74V859l-421 421h330l91-91zm-603 91l512-512h-330l-512 512h330zm-512 0l512-512H987l-512 512h330zM128 640h1792V128H128v512z"/>`),
		g.Group(children),
	)
}