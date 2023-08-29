package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#75B2DD" d="M.5.552h300v157.895H.5z"/><path fill="#FFF" d="m135.483 27.252l9.281 6.743l-3.545 10.91l9.281-6.743l9.281 6.743l-3.545-10.91l9.281-6.743h-11.472l-3.545-10.91l-3.493 10.91zM87.342 79.5l10.91 3.545v11.472l6.743-9.281l10.911 3.545l-6.743-9.281l6.743-9.281l-10.911 3.545l-6.743-9.281v11.472zm48.141 52.248h11.472l3.545 10.91l3.545-10.91h11.472l-9.281-6.743l3.545-10.911l-9.281 6.743l-9.281-6.743l3.545 10.911zm49.611-61.529l6.744 9.281l-6.744 9.281l10.911-3.545l6.743 9.281V83.045l10.91-3.545l-10.91-3.545V64.483l-6.743 9.281z"/></g>`),
		g.Group(children),
	)
}