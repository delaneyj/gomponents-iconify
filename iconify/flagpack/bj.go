package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bj(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackBj0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackBj1)"><use href="#flagpackBj0"/><path fill="#DD2C2B" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="M0 0v14h32V0H0Z" clip-rule="evenodd"/><path fill="#5EAA22" d="M0 0h14v24H0z"/></g><defs><clipPath id="flagpackBj1"><use href="#flagpackBj0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}