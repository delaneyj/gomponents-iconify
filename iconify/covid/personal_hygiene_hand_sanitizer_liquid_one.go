package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalHygieneHandSanitizerLiquidOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m1.121 19.788l2.905 2.324a5.187 5.187 0 0 0 3.243 1.138h5.967a1.731 1.731 0 0 0 1.731-1.731v0a1.731 1.731 0 0 0-1.731-1.731H8.765"/><path d="M1.278 14.6H4.74c1.026 0 2.029.304 2.882.873L9.6 16.788a1.615 1.615 0 0 1 .513 2.281v0a1.614 1.614 0 0 1-2.066.549l-2.6-1.56m17.432-7.418a4.945 4.945 0 1 1-9.89 0c0-3.709 4.945-9.89 4.945-9.89s4.945 6.181 4.945 9.89Zm-6.923-.989h3.956m-1.978-1.978v3.956"/></g>`),
		g.Group(children),
	)
}