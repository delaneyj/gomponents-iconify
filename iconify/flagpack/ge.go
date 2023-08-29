package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGe0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGe2)"><use href="#flagpackGe0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackGe1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill="#E31D1C" fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackGe1)"><path d="M14 0h4v10h14v4H18v10h-4V14H0v-4h14V0Z"/><path d="M9.999 17.222L12.2 17v2s-2.201-.138-2.201-.098c0 .04.201 2.098.201 2.098h-2l.16-2H6.2v-2l2.16.222L8.2 15h2l-.201 2.222Zm0-12L12.2 5v2s-2.201-.138-2.201-.098c0 .04.201 2.098.201 2.098h-2l.16-2H6.2V5l2.16.222L8.2 3h2l-.201 2.222Zm14 0L26.2 5v2s-2.201-.138-2.201-.098c0 .04.201 2.098.201 2.098h-2l.16-2H20.2V5l2.16.222L22.2 3h2l-.201 2.222Zm0 12L26.2 17v2s-2.201-.138-2.201-.098c0 .04.201 2.098.201 2.098h-2l.16-2H20.2v-2l2.16.222L22.2 15h2l-.201 2.222Z"/></g></g><defs><clipPath id="flagpackGe2"><use href="#flagpackGe0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}