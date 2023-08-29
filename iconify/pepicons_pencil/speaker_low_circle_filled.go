package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerLowCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilSpeakerLowCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M12.672 5.123L8.399 8.84H5a.5.5 0 0 0-.5.5v5.2a.5.5 0 0 0 .5.5h3.295l4.375 3.835a.5.5 0 0 0 .83-.376v-13a.5.5 0 0 0-.828-.377ZM8.884 9.745L12.5 6.598v10.799L9.014 14.34a.5.5 0 0 0-.458-.3H5.5v-4.2h2.894a.492.492 0 0 0 .49-.096Z" clip-rule="evenodd"/><path d="M16.326 14.88a.5.5 0 0 1-.652-.76c.04-.033.079-.07.117-.11c.157-.162.295-.366.407-.602c.195-.409.302-.896.302-1.408c0-.817-.273-1.558-.709-2.01a1.738 1.738 0 0 0-.117-.11a.5.5 0 0 1 .652-.76c.064.056.125.114.185.176c.624.647.989 1.639.989 2.704c0 .66-.14 1.293-.398 1.838a3.28 3.28 0 0 1-.591.866a2.76 2.76 0 0 1-.185.175Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilSpeakerLowCircleFilled0)"/></g>`),
		g.Group(children),
	)
}