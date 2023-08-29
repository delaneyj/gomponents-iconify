package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhostLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M110 116a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm46-10a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm66 14v96a6 6 0 0 1-9.8 4.64l-25.53-20.89l-25.54 20.89a6 6 0 0 1-7.6 0L128 199.75l-25.53 20.89a6 6 0 0 1-7.6 0l-25.54-20.89l-25.53 20.89A6 6 0 0 1 34 216v-96a94 94 0 0 1 188 0Zm-12 0a82 82 0 0 0-164 0v83.34l19.53-16a6 6 0 0 1 7.6 0l25.54 20.89l25.53-20.89a6 6 0 0 1 7.6 0l25.53 20.89l25.54-20.89a6 6 0 0 1 7.6 0l19.53 16Z"/>`),
		g.Group(children),
	)
}