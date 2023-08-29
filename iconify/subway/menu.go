package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M186.2 139.6h139.6V0H186.2v139.6zM372.4 0v139.6H512V0H372.4zM0 139.6h139.6V0H0v139.6zm186.2 186.2h139.6V186.2H186.2v139.6zm186.2 0H512V186.2H372.4v139.6zM0 325.8h139.6V186.2H0v139.6zM186.2 512h139.6V372.4H186.2V512zm186.2 0H512V372.4H372.4V512zM0 512h139.6V372.4H0V512z"/>`),
		g.Group(children),
	)
}