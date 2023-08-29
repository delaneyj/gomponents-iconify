package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cola(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="m36 37l-2.79 6.4a1 1 0 0 1-.918.6H15.707a1 1 0 0 1-.916-.6L12 37m0-26l2.79-6.4a1 1 0 0 1 .918-.6h16.584a1 1 0 0 1 .917.6L36 11"/><path d="M12 12a2 2 0 0 1 2-2h20a2 2 0 0 1 2 2v24a2 2 0 0 1-2 2H14a2 2 0 0 1-2-2V12Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m13 22.5l6.396-1.01a8.934 8.934 0 0 0 5.218-2.815v0a8.935 8.935 0 0 1 7.168-2.907L35.5 16m-23 16.5l5.553-.252a9.54 9.54 0 0 0 7.998-5.067v0a9.54 9.54 0 0 1 5.974-4.754L35.5 21.5M36 12v20M12 16v20"/></g>`),
		g.Group(children),
	)
}