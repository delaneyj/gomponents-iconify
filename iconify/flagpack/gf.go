package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGf0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGf1)"><use href="#flagpackGf0"/><rect width="32" height="24" fill="#5EAA22" rx="2"/><path fill="#FECA00" fill-rule="evenodd" d="m0 0l32 24H0V0Z" clip-rule="evenodd"/><path fill="#E21835" fill-rule="evenodd" d="m15.93 14.406l-3.485 2.418l1.114-4.141L11 10.038l3.465-.143L15.93 5.8l1.466 4.095h3.458l-2.553 2.788l1.279 3.897l-3.65-2.174Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackGf1"><use href="#flagpackGf0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}