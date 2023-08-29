package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGb0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGb1)"><use href="#flagpackGb0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><path fill="#F50302" fill-rule="evenodd" d="M18 0h-4v10H0v4h14v10h4V14h14v-4H18V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackGb1"><use href="#flagpackGb0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}