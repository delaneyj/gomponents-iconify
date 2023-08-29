package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 13.984a.499.499 0 0 1-.479-.358L6.647 8.984H4.993a.5.5 0 0 1-.486-.383l-.976-4.052l-1.048 4.06a.501.501 0 0 1-.484.375H0v-.953h1.61l1.452-5.625a.5.5 0 0 1 .484-.375h.004a.5.5 0 0 1 .482.383l1.352 5.617h1.635c.222 0 .417.146.479.358l1.005 3.346l1.02-3.348a.499.499 0 0 1 .479-.356h.687l1.347-2.736a.501.501 0 0 1 .445-.279h.004c.188 0 .359.104.444.271l1.41 2.744h1.63v.953h-1.936a.496.496 0 0 1-.444-.271l-1.095-2.131l-1.047 2.123a.497.497 0 0 1-.447.279h-.626l-1.396 4.644a.497.497 0 0 1-.478.356z"/>`),
		g.Group(children),
	)
}