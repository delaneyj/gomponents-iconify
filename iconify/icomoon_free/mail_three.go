package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.333 0H2.667A2.675 2.675 0 0 0 0 2.667v10.666C0 14.801 1.2 16 2.667 16h10.666A2.674 2.674 0 0 0 16 13.333V2.667C16 1.2 14.8 0 13.333 0zM2.854 13.854l-1.207-1.207l4-4l.457.457l-3.25 4.75zm-.458-10.75l.457-.457l5.146 4.146l5.146-4.146l.457.457l-5.604 6.604l-5.604-6.604zm10.75 10.75l-3.25-4.75l.457-.457l4 4l-1.207 1.207z"/>`),
		g.Group(children),
	)
}