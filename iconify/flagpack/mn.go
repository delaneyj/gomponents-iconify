package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackMn0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackMn1)"><use href="#flagpackMn0"/><path fill="#4C67E8" fill-rule="evenodd" d="M10 0h12v24H10V0Z" clip-rule="evenodd"/><path fill="#C51918" fill-rule="evenodd" d="M22 0h10v24H22V0ZM0 0h12v24H0V0Z" clip-rule="evenodd"/><path fill="#F8D000" fill-rule="evenodd" d="M6.014 6.884c-1.013 0-.88-.987-.88-.987l.48-1.038v1.038c0 .133.182-.134.182-.455s.218-.802.218-.802s.009-.161.015-.32c.034.082.104.152.21.26l.06.06c.124.128.108.468.093.762c-.013.266-.025.495.073.495c.206 0 .19-.919.19-.919l.346.919s.026.987-.987.987Zm0-2.82c.026-.15.022.051.015.256a.435.435 0 0 1-.015-.255Zm1.524 4.62c0 .802-.671 1.453-1.5 1.453c-.828 0-1.5-.65-1.5-1.453s.672-1.454 1.5-1.454c.829 0 1.5.651 1.5 1.454Zm-3.769 2.72H2v8.572h1.77v-8.572Zm6.231 0H8.23v8.572H10v-8.572Zm-5.77.075l1.725 1.068L7.9 11.479H4.231ZM5.956 20l-1.724-1.068H7.9L5.955 20Zm-1.724-7.105h3.615v.67H4.231v-.67Zm3.615 4.994H4.231v.67h3.615v-.67Zm-1.77-.373c1.02 0 1.847-.8 1.847-1.789c0-.988-.826-1.789-1.846-1.789s-1.846.801-1.846 1.79c0 .987.826 1.788 1.846 1.788Zm-2.073-8.3s.062 1.868 1.903 1.868c1.84 0 2.137-1.868 2.137-1.868s-.712 1.144-2.02 1.144c-1.307 0-2.02-1.144-2.02-1.144Z" clip-rule="evenodd"/><path fill="#C51918" d="m6.359 13.972l.082-.367c.412.451.6.704.511 1.118c-.098.461-.617.975-1.219 1.399c-.364.256-.587.949-.406 1.169l-.173.314c-.403-.49-.081-1.417.475-1.808c.511-.36.868-.737.93-1.033c.04-.194.117-.445-.2-.792Z"/><path fill="#C51918" fill-rule="evenodd" d="M6.5 17a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-1.124-1.35a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackMn1"><use href="#flagpackMn0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}