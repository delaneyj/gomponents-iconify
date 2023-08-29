package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm0-6a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd" opacity=".5"/><path d="m5.479 19.582l4.272-4.273a4.021 4.021 0 0 1-1.06-1.06L4.418 18.52c.327.38.681.734 1.06 1.06ZM4.418 5.48l4.273 4.272a4.022 4.022 0 0 1 1.06-1.06L5.48 4.417c-.38.327-.734.682-1.061 1.061Zm9.83 3.211L18.52 4.42c.379.326.734.68 1.06 1.06l-4.27 4.272a4.021 4.021 0 0 0-1.061-1.06Zm5.333 9.83l-4.273-4.273a4.02 4.02 0 0 1-1.06 1.06l4.272 4.274c.38-.327.735-.682 1.061-1.061Z"/></g>`),
		g.Group(children),
	)
}