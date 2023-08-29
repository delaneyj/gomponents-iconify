package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextClearFormattingGaTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M12.5 3.25a.75.75 0 0 0-1.5 0v8.5a.75.75 0 0 0 1.5 0V8H14a.75.75 0 0 0 0-1.5h-1.5V3.25zm-8.75.25a.75.75 0 1 0 0 1.5h4.245c-.162 1.634-1.328 4.092-4.46 5.032a.75.75 0 0 0 .43 1.436c4.39-1.317 5.726-5.15 5.532-7.286A.75.75 0 0 0 8.75 3.5h-5zm16.252 18h-3.02l5.446-5.447a1.95 1.95 0 0 0 .002-2.759l-2.724-2.723a1.95 1.95 0 0 0-2.759.001l-6.374 6.375a1.95 1.95 0 0 0-.002 2.759l2.724 2.724a1.94 1.94 0 0 0 1.208.563L14.5 23H20a.75.75 0 0 0 0-1.5zm-6.151-5.71l4.157-4.157a.45.45 0 0 1 .637-.002l2.724 2.724a.45.45 0 0 1-.001.638l-4.157 4.156l-3.36-3.36zm-2.218 2.218l1.157-1.157l3.36 3.36l-1.157 1.156a.45.45 0 0 1-.637.002l-2.724-2.724a.45.45 0 0 1 .001-.637z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}