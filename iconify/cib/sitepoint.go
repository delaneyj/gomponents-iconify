package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sitepoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m3.297 14.042l2.359 2.255l7.464 6.854l3.198-3.057c.281-.396.26-.938-.057-1.313l-2.943-2.536l.016-.005l-3.161-3.021a1.054 1.054 0 0 1 .026-1.438l8.521-8.104L14.861 0L3.304 10.984a2.092 2.092 0 0 0 0 3.057zm25.406 3.916l-2.359-2.255l-7.464-6.854l-3.214 3.057a1.03 1.03 0 0 0 .063 1.313l2.938 2.521h-.005l3.156 3.021c.359.417.359 1.021-.036 1.417l-8.521 8.099l3.88 3.724l11.557-10.984a2.087 2.087 0 0 0 0-3.057z"/>`),
		g.Group(children),
	)
}