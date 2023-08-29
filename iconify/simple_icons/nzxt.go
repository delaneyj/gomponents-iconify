package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nzxt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m1.763 8.936l2.101 3.04l-.002-3.04h1.773v6.128H3.99l-2.218-3.227v3.227H0V8.936zm22.237 0v1.614h-1.612v4.514h-1.883V10.55h-1.611V8.936H24zm-9.598 0l.996 1.573l1.061-1.573h1.947l-1.98 3.034l2.013 3.094h-2.063l-1.005-1.558l-.99 1.558h-2.015l1.975-3.038l-2.004-3.09h2.065zm-2.652 0L9.327 13.51h2.372v1.554H6.573l2.379-4.584H6.704V8.936h5.046z"/>`),
		g.Group(children),
	)
}