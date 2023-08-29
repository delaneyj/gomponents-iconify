package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudFogLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M120 198H72a6 6 0 0 1 0-12h48a6 6 0 0 1 0 12Zm64-12h-24a6 6 0 0 0 0 12h24a6 6 0 0 0 0-12Zm-24 32h-56a6 6 0 0 0 0 12h56a6 6 0 0 0 0-12Zm70-126a74.09 74.09 0 0 1-74 74H76a50 50 0 1 1 10.2-99A74.08 74.08 0 0 1 230 92Zm-12 0a62.06 62.06 0 0 0-124-3.65a6 6 0 0 1-12-.7a75.84 75.84 0 0 1 1.07-9A38 38 0 1 0 76 154h80a62.07 62.07 0 0 0 62-62Z"/>`),
		g.Group(children),
	)
}