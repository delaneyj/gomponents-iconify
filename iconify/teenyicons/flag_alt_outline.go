package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m12.5 6.5l.224.447a.5.5 0 0 0 .033-.876L12.5 6.5Zm-10-6l.257-.429A.5.5 0 0 0 2 .5h.5Zm10.257 5.571l-10-6l-.514.858l10 6l.514-.858ZM2 .5v11h1V.5H2Zm.724 11.447l10-5l-.448-.894l-10 5l.448.894ZM3 15v-3.5H2V15h1Z"/>`),
		g.Group(children),
	)
}