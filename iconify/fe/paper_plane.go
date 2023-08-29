package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePaperPlane0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePaperPlane1" fill="currentColor" fill-rule="nonzero"><path id="fePaperPlane2" d="m13.761 12.01l-10.76-1.084L3 4.074a1.074 1.074 0 0 1 1.554-.96l15.852 7.926a1.074 1.074 0 0 1 0 1.92l-15.85 7.926a1.074 1.074 0 0 1-1.554-.96v-6.852l10.76-1.064Z"/></g></g>`),
		g.Group(children),
	)
}