package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Invert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 48h768v598H0V48zm384 547v-71l177-177l-177-177V99H51v496h333zm0-71L207 347l177-177v354z"/>`),
		g.Group(children),
	)
}