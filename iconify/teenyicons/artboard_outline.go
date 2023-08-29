package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArtboardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4.5 4.5V4a.5.5 0 0 0-.5.5h.5Zm6 0h.5a.5.5 0 0 0-.5-.5v.5Zm0 6v.5a.5.5 0 0 0 .5-.5h-.5Zm-6 0H4a.5.5 0 0 0 .5.5v-.5ZM4 0v2h1V0H4Zm6 0v2h1V0h-1ZM0 5h2V4H0v1Zm0 6h2v-1H0v1Zm13-6h2V4h-2v1Zm0 6h2v-1h-2v1Zm-9 2v2h1v-2H4Zm6 0v2h1v-2h-1ZM4.5 5h6V4h-6v1Zm5.5-.5v6h1v-6h-1Zm.5 5.5h-6v1h6v-1Zm-5.5.5v-6H4v6h1Z"/>`),
		g.Group(children),
	)
}