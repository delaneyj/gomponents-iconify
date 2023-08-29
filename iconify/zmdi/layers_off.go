package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m380 304l-31-31l26-19l30 30zm-10-101l-51 40L151 75l62-48l192 149zM27 5l400 400l-27 27l-81-81l-106 82L21 284l35-27l157 123l76-59l-31-30l-45 34L56 203l-35-27l69-54L0 32z"/>`),
		g.Group(children),
	)
}