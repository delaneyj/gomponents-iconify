package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighBarsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M32 9a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v30a3 3 0 0 1-3 3h-4a3 3 0 0 1-3-3V9Zm3-1a1 1 0 0 0-1 1v1.172L36.172 8H35Zm4 0l-5 5v2.172l6-6V9a1 1 0 0 0-1-1Zm1 4l-6 6v2.172l6-6V12Zm0 5l-6 6v2.172l6-6V17Zm0 5l-6 6v2.172l6-6V22Zm0 5l-6 6v2.172l6-6V27Zm0 5l-6 6v1c0 .32.15.605.384.788L40 34.172V32Zm0 5l-3 3h2a1 1 0 0 0 1-1v-2ZM19 21a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v18a3 3 0 0 1-3 3h-4a3 3 0 0 1-3-3V21Zm5-1h2a.93.93 0 0 1 .159.013L21 25.172V23l3-3Zm-3 10.172V28l6-6v2.172l-6 6Zm6-1V27l-6 6v2.172l6-6ZM21 39v-1l6-6v2.172l-5.616 5.616A.998.998 0 0 1 21 39Zm3 1l3-3v2a1 1 0 0 1-1 1h-2ZM9 30a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3v-6a3 3 0 0 0-3-3H9Zm-1 3a1 1 0 0 1 1-1h2.172L8 35.172V33Zm0 5l5.707-5.707A.997.997 0 0 1 14 33v1.172l-5.616 5.616A.998.998 0 0 1 8 39v-1Zm3 2h2a1 1 0 0 0 1-1v-2l-3 3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}