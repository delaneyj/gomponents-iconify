package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><path d="M12.971 15.23c0 .423-.417.768-.932.768H3.951c-.515 0-.932-.345-.932-.768v-.45c0-.422.417-.767.932-.767h8.088c.515 0 .932.345.932.767v.45zm-.506-2.251H3.547L1.421 5.365l1.141-.709l3.011 4.251l1.772-5.865h1.418l1.594 6.22l3.188-4.606l1.046.709l-2.126 7.614z"/><circle cx="14.968" cy="2.968" r=".968"/><circle cx="7.964" cy=".964" r=".964"/><ellipse cx=".983" cy="2.974" rx=".983" ry=".974"/></g>`),
		g.Group(children),
	)
}