package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Remove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 264.84L335.16 600L0 935.16L264.84 1200L600 864.84L935.16 1200L1200 935.16L864.84 600L1200 264.84L935.16 0L600 335.16L264.84 0L0 264.84z"/>`),
		g.Group(children),
	)
}