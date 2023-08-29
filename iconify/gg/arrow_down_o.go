package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m14.829 12.026l1.414 1.414L12 17.683L7.757 13.44l1.415-1.414L11 13.854V6.317h2v7.537l1.829-1.828Z"/><path fill-rule="evenodd" d="M19.778 19.778c-4.296 4.296-11.26 4.296-15.556 0c-4.296-4.296-4.296-11.26 0-15.556c4.296-4.296 11.26-4.296 15.556 0c4.296 4.296 4.296 11.26 0 15.556Zm-1.414-1.414A9 9 0 1 1 5.636 5.636a9 9 0 0 1 12.728 12.728Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}