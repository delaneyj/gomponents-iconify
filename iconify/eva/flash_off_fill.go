package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaFlashOffFill0"><g id="evaFlashOffFill1"><path id="evaFlashOffFill2" fill="currentColor" d="m17.33 14.5l2.5-3.74A1 1 0 0 0 19 9.2h-5.89l.77-7.09a1 1 0 0 0-.65-1a1 1 0 0 0-1.17.38L8.94 6.11Zm-10.66-5l-2.5 3.74A1 1 0 0 0 5 14.8h5.89l-.77 7.09a1 1 0 0 0 .65 1.05a1 1 0 0 0 .34.06a1 1 0 0 0 .83-.44l3.12-4.67Zm14.04 9.79l-16-16a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g>`),
		g.Group(children),
	)
}