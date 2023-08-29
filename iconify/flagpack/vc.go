package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackVc0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackVc1)"><use href="#flagpackVc0"/><path fill="#FFDC17" fill-rule="evenodd" d="M8 0h16v24H8V0Z" clip-rule="evenodd"/><path fill="#5FBF2B" fill-rule="evenodd" d="M24 0h8v24h-8V0Z" clip-rule="evenodd"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0h8v24H0V0Z" clip-rule="evenodd"/><path fill="#5FBF2B" fill-rule="evenodd" d="m19.726 5.6l-3.04 4.463l3.04 4.305l3.04-4.305l-3.04-4.463Zm-7.449.079L9 10.063l3.04 4.226l3.04-4.226l-2.803-4.384Zm.566 9.31l3.277-4.384l2.803 4.384L16.12 19.5l-3.277-4.512Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackVc1"><use href="#flagpackVc0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}