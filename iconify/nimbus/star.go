package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m8 3.12l1.43 2.77l.31.59l.66.08l3.12.35L11.26 9l-.5.47l.12.67l.53 2.95l-2.85-1.42l-.56-.28l-.57.28l-2.85 1.44l.53-2.95l.12-.67l-.5-.49l-2.26-2.08l3.12-.36l.66-.08l.31-.59L8 3.12m0-2a.63.63 0 0 0-.57.33l-2 3.84l-4.51.55a.59.59 0 0 0-.34 1l3.3 3.08l-.76 4.21a.61.61 0 0 0 .62.7a.65.65 0 0 0 .26-.05l4-2l4 2a.65.65 0 0 0 .29.07a.61.61 0 0 0 .62-.7l-.76-4.21l3.31-3.09a.59.59 0 0 0-.34-1l-4.54-.51l-2-3.84A.63.63 0 0 0 8 1.15z"/>`),
		g.Group(children),
	)
}