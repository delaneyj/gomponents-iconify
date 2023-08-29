package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soundright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m960 575l54 135q9 10 9 23.5t-9.5 23.5t-23.5 10t-24-10l-70-182h-64v160q0 14-9.5 23t-22.5 9t-22.5-9t-9.5-23V288q0-14 9.5-23.5T800 255h96q53 0 90.5 37.5T1024 383v65q0 38-18.5 73T960 575zm0-192q0-26-19-45t-45-19h-64v192h64q26 0 45-18.5t19-45.5v-64zm-395 628L324 768h-4V255h4L565 12q30-30 75 16v968q-45 46-75 15zM0 704V319q0-26 18.5-45T64 255h192v513H64q-27 0-45.5-19T0 704z"/>`),
		g.Group(children),
	)
}