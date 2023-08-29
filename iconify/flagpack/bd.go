package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackBd0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackBd1)"><use href="#flagpackBd0"/><path fill="#38A17E" d="M0 0h32v24H0z"/><path fill="#F72E45" fill-rule="evenodd" d="M13 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackBd1"><use href="#flagpackBd0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}