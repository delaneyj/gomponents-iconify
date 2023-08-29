package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Observatory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 724v128q0 27-19 45.5T896 916H128q-27 0-45.5-18.5T64 852V724q-27 0-45.5-18.5T0 660t18.5-45.5T64 596V404q0-132 70.5-240.5T320 0v404q0 27 18.5 45.5T384 468V148q0-53 37.5-90.5T512 20t90.5 37.5T640 148v320q27 0 45.5-18.5T704 404V0q115 55 185.5 163.5T960 404v192q27 0 45.5 18.5t18.5 45t-19 45.5t-45 19zM512 84q-27 0-45.5 19T448 148.5t18.5 45T512 212t45.5-18.5t18.5-45t-18.5-45.5T512 84z"/>`),
		g.Group(children),
	)
}