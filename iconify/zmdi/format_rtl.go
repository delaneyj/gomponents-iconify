package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 173q-35 0-60-25T43 88t25-60t60-25h171v42h-43v235h-43V45h-42v235h-43V173zM85 323h256v42H85v64L0 344l85-85v64z"/>`),
		g.Group(children),
	)
}