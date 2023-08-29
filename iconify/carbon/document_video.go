package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m30 19l-4 3.2V20a2.002 2.002 0 0 0-2-2h-8a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h8a2.002 2.002 0 0 0 2-2v-2.2l4 3.2Zm-14 7v-6h8l.002 6Z"/><path fill="currentColor" d="M12 28H8V4h8v6a2.006 2.006 0 0 0 2 2h6v3h2v-5a.91.91 0 0 0-.3-.7l-7-7A.909.909 0 0 0 18 2H8a2.006 2.006 0 0 0-2 2v24a2.006 2.006 0 0 0 2 2h4Zm6-23.6l5.6 5.6H18Z"/>`),
		g.Group(children),
	)
}