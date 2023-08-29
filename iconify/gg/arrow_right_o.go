package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m12.052 14.829l1.414 1.414L17.71 12l-4.243-4.243l-1.414 1.415L13.88 11H6.343v2h7.537l-1.828 1.829Z"/><path fill-rule="evenodd" d="M19.778 19.778c4.296-4.296 4.296-11.26 0-15.556c-4.296-4.296-11.26-4.296-15.556 0c-4.296 4.296-4.296 11.26 0 15.556c4.296 4.296 11.26 4.296 15.556 0Zm-1.414-1.414A9 9 0 1 0 5.636 5.636a9 9 0 0 0 12.728 12.728Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}