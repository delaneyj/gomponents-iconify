package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func So(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackSo0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackSo1)"><use href="#flagpackSo0"/><path fill="#56C6F5" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="m16.179 14.717l-3.764 2.393l1.262-4.189l-2.684-2.737l3.701-.08l1.637-4.137l1.493 4.19l3.692.065l-2.775 2.788l1.296 3.985l-3.858-2.278Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackSo1"><use href="#flagpackSo0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}