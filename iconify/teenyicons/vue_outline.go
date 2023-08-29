package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VueOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5 13.5l-.432.252a.5.5 0 0 0 .864 0L7.5 13.5Zm4-12l.434.248A.5.5 0 0 0 11.5 1v.5Zm-4 7l-.434.248a.5.5 0 0 0 .868 0L7.5 8.5Zm-4-7V1a.5.5 0 0 0-.434.748L3.5 1.5Zm3 0l.447-.224L6.81 1H6.5v.5Zm1 2l-.447.224a.5.5 0 0 0 .894 0L7.5 3.5Zm1-2V1h-.309l-.138.276l.447.224Zm-8.432.252l7 12l.864-.504l-7-12l-.864.504Zm7.864 12l7-12l-.864-.504l-7 12l.864.504Zm3.134-12.5l-4 7l.868.496l4-7l-.868-.496Zm-3.132 7l-4-7l-.868.496l4 7l.868-.496ZM3.5 2h3V1h-3v1Zm2.553-.276l1 2l.894-.448l-1-2l-.894.448Zm1.894 2l1-2l-.894-.448l-1 2l.894.448ZM8.5 2h3V1h-3v1Z"/>`),
		g.Group(children),
	)
}