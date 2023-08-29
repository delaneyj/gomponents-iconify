package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m14.779 12.18l-2.984-3.679s1.601.436 1.775.436c.465 0 .195-.517 0-.714l-2.385-3.031s1.148-.274 1.565-.274c.418 0 .197-.517 0-.714L9.4.061a.505.505 0 0 0-.714 0l-3.395 4.1c-.198.197-.486.715 0 .715s1.622.316 1.622.316L4.325 8.079c-.198.197-.557.714 0 .714c.319 0 1.95-.291 1.95-.291l-2.958 3.687c-.197.196-.557.714 0 .714s4.691-1.007 4.691-1.007v3.045c0 .537.436.973.975.973c.537 0 1.015-.436 1.015-.973v-3.045s4.375.999 4.78.999c.405 0 .198-.519.001-.715z"/>`),
		g.Group(children),
	)
}