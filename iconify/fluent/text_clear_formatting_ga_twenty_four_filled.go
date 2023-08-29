package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextClearFormattingGaTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M12.5 3.25a.75.75 0 0 0-1.5 0v8.5a.75.75 0 0 0 1.5 0V8H14a.75.75 0 0 0 0-1.5h-1.5V3.25zm-8.75.25a.75.75 0 1 0 0 1.5h4.245c-.162 1.634-1.328 4.092-4.46 5.032a.75.75 0 0 0 .43 1.436c4.39-1.317 5.726-5.15 5.532-7.286A.75.75 0 0 0 8.75 3.5h-5zm14.522 16.711L12.79 14.73l4.156-4.157a1.952 1.952 0 0 1 2.759-.001l2.724 2.724a1.95 1.95 0 0 1-.002 2.759l-4.156 4.156zm-6.542-4.42l5.481 5.48l-.23.23h3.02a.75.75 0 0 1 0 1.5h-5.5l.002-.007a1.94 1.94 0 0 1-1.208-.563l-2.724-2.724a1.95 1.95 0 0 1 .002-2.759l1.157-1.157z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}