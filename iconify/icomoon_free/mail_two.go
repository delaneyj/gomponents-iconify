package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.333 0H2.667A2.675 2.675 0 0 0 0 2.667v10.666C0 14.801 1.2 16 2.667 16h10.666A2.674 2.674 0 0 0 16 13.333V2.667C16 1.2 14.8 0 13.333 0zm0 2c.125 0 .243.036.344.099L7.999 6.793L2.322 2.099A.648.648 0 0 1 2.666 2h10.666zM2.667 14a.654.654 0 0 1-.089-.006l3.525-4.89l-.457-.457L2 12.293V2.744L8 10l6-7.256v9.549l-3.646-3.646l-.457.457l3.525 4.89a.65.65 0 0 1-.088.006H2.668z"/>`),
		g.Group(children),
	)
}