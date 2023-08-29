package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSnowLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M86 196a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm30 6a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm48-16a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm-96 40a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm88 0a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm74-134a74.09 74.09 0 0 1-74 74H76a50 50 0 1 1 10.2-99A74.08 74.08 0 0 1 230 92Zm-12 0a62.06 62.06 0 0 0-124-3.65a6 6 0 0 1-12-.7a75.84 75.84 0 0 1 1.07-9A38 38 0 1 0 76 154h80a62.07 62.07 0 0 0 62-62Z"/>`),
		g.Group(children),
	)
}