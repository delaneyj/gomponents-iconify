package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.526 6.106c.5.233.74.777.537 1.215L9.884 18.424c-.204.438-.775.604-1.276.37c-.5-.233-.74-.778-.536-1.216L13.25 6.476c.204-.438.775-.604 1.276-.37Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}