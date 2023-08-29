package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Suitcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 107h-85V64q0-27-18.5-45.5T299 0h-86q-27 0-45.5 18.5T149 64v43H64q-27 0-45.5 18T0 171v277q0 27 18.5 45.5T64 512h384q27 0 45.5-18.5T512 448V171q0-28-18.5-46T448 107zM192 64q0-21 21-21h86q21 0 21 21v43H192V64zM64 469q-21 0-21-21V171q0-22 21-22h21v320H64zm64 0V149h256v320H128zm341-21q0 21-21 21h-21V149h21q21 0 21 22v277z"/>`),
		g.Group(children),
	)
}