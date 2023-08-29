package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M9.672 2.123L5.399 5.84H2a.5.5 0 0 0-.5.5v5.2a.5.5 0 0 0 .5.5h3.295l4.375 3.835a.5.5 0 0 0 .83-.376v-13a.5.5 0 0 0-.828-.377ZM5.884 6.745L9.5 3.598v10.799L6.014 11.34a.5.5 0 0 0-.458-.3H2.5v-4.2h2.894a.492.492 0 0 0 .49-.096Z" clip-rule="evenodd"/><path d="M13.326 11.88a.5.5 0 0 1-.652-.76c.04-.033.079-.07.117-.11c.157-.162.295-.366.407-.602c.195-.409.302-.896.302-1.408c0-.817-.273-1.558-.709-2.01a1.738 1.738 0 0 0-.117-.11a.5.5 0 0 1 .652-.76c.064.056.125.114.185.176c.624.647.989 1.639.989 2.704c0 .66-.14 1.293-.399 1.838c-.157.33-.356.624-.59.866a2.76 2.76 0 0 1-.185.175Z"/></g>`),
		g.Group(children),
	)
}