package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Powerdirector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.514 35.107H4.5V12.893h30.014m-30.014 0l30.014 22.214m0-22.214L4.5 35.107"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.692 24l7.822 5.358V18.643L26.692 24zM43.5 34.733V13.267m-4.493 17.659V17.074"/>`),
		g.Group(children),
	)
}