package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boorusphere(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.038 42.5c.183-16.554 9.914-20.46 17.556-22.449c14.448-3.76 21.064 13.1 5.113 22.449H8.038Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.038 42.5V18.34c0-7.656 7.03-12.912 13.873-12.84c11.08.119 9.326 9.055 9.77 14.183"/>`),
		g.Group(children),
	)
}