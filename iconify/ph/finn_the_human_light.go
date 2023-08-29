package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FinnTheHumanLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 98H88a38 38 0 0 0-38 38v8a38 38 0 0 0 38 38h80a38 38 0 0 0 38-38v-8a38 38 0 0 0-38-38Zm26 46a26 26 0 0 1-26 26H88a26 26 0 0 1-26-26v-8a26 26 0 0 1 26-26h80a26 26 0 0 1 26 26Zm14-110a30.05 30.05 0 0 0-29.4 24H77.4A30 30 0 0 0 18 64v80a70.08 70.08 0 0 0 70 70h80a70.08 70.08 0 0 0 70-70V64a30 30 0 0 0-30-30Zm18 110a58.07 58.07 0 0 1-58 58H88a58.07 58.07 0 0 1-58-58V64a18 18 0 0 1 36 0a6 6 0 0 0 6 6h112a6 6 0 0 0 6-6a18 18 0 0 1 36 0Zm-124-4a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm72 0a10 10 0 1 1-10-10a10 10 0 0 1 10 10Z"/>`),
		g.Group(children),
	)
}