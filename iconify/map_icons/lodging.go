package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lodging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M46 32v8h4V28H4V11.66c0-1.123-.869-2.042-2-2.042c-1.127 0-2 .918-2 2.042V40h4v-8h42zM10.33 19.852a3.687 3.687 0 0 0 3.687-3.694a3.68 3.68 0 0 0-3.687-3.675a3.678 3.678 0 0 0-3.683 3.675a3.686 3.686 0 0 0 3.683 3.694zM50 25l-.014-4.871c-.013-1.606-1.36-2.618-2.939-2.809L18.09 14.487l-.215-.009c-1.006 0-1.875.834-1.875 1.835V22H8.547c-1.011 0-1.826.5-1.826 1.499c0 1.015.815 1.501 1.826 1.501H50z"/>`),
		g.Group(children),
	)
}