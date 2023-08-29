package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackLc0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackLc1)"><use href="#flagpackLc0"/><path fill="#7CCFF5" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="m16 4l8 16H8l8-16Z" clip-rule="evenodd"/><path fill="#272727" fill-rule="evenodd" d="m16 8l7 12H9l7-12Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="m16 14l8 6H8l8-6Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackLc1"><use href="#flagpackLc0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}