package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Erase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 12.9a2 2 0 0 0 0 2.828l3.858 3.858H4.086a1 1 0 1 0 0 2h16a1 1 0 0 0 0-2h-9.13l9.515-9.515a2 2 0 0 0 0-2.828L16.227 3a2 2 0 0 0-2.829 0L3.5 12.9Zm4.326-1.498l-2.912 2.912l4.243 4.242l2.911-2.911l-4.242-4.243ZM9.24 9.988l4.243 4.242l5.573-5.573l-4.242-4.243L9.24 9.988Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}