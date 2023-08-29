package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chevrons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M736 1024H576l288-512L576 0h160l288 512zm-448 0L0 512L288 0h160L160 512l288 512H288z"/>`),
		g.Group(children),
	)
}