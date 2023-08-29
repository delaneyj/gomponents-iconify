package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaFlashOutline0"><g id="evaFlashOutline1"><path id="evaFlashOutline2" fill="currentColor" d="M11.11 23a1 1 0 0 1-.34-.06a1 1 0 0 1-.65-1.05l.77-7.09H5a1 1 0 0 1-.83-1.56l7.89-11.8a1 1 0 0 1 1.17-.38a1 1 0 0 1 .65 1l-.77 7.14H19a1 1 0 0 1 .83 1.56l-7.89 11.8a1 1 0 0 1-.83.44ZM6.87 12.8H12a1 1 0 0 1 .74.33a1 1 0 0 1 .25.78l-.45 4.15l4.59-6.86H12a1 1 0 0 1-1-1.11l.45-4.15Z"/></g></g>`),
		g.Group(children),
	)
}