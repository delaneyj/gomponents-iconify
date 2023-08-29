package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m11.948 14.829l-1.414 1.414L6.29 12l4.243-4.243l1.414 1.415L10.12 11h7.537v2H10.12l1.828 1.829Z"/><path fill-rule="evenodd" d="M4.222 19.778c-4.296-4.296-4.296-11.26 0-15.556c4.296-4.296 11.26-4.296 15.556 0c4.296 4.296 4.296 11.26 0 15.556c-4.296 4.296-11.26 4.296-15.556 0Zm1.414-1.414A9 9 0 1 1 18.364 5.636A9 9 0 0 1 5.636 18.364Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}