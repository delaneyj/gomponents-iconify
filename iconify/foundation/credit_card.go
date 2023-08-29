package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M86.941 26.224a2.66 2.66 0 0 0-2.66-2.659c-.036 0-.07.009-.106.011H15.949c-.078-.007-.153-.023-.233-.023a2.659 2.659 0 0 0-2.658 2.659c0 .124.02.243.037.363v7.009h73.846v-7.243h-.011c.001-.041.011-.078.011-.117zM13.095 73.78v.01a2.659 2.659 0 0 0 2.659 2.658c.056 0 .109-.013.164-.017v.002h68.459v-.02a2.653 2.653 0 0 0 2.563-2.633h.002V43.582H13.095V73.78z"/>`),
		g.Group(children),
	)
}