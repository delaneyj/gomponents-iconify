package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnifyingGlassMinusThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M148 112a4 4 0 0 1-4 4H80a4 4 0 0 1 0-8h64a4 4 0 0 1 4 4Zm78.83 114.83a4 4 0 0 1-5.66 0l-52.7-52.7a84.1 84.1 0 1 1 5.66-5.66l52.7 52.7a4 4 0 0 1 0 5.66ZM112 188a76 76 0 1 0-76-76a76.08 76.08 0 0 0 76 76Z"/>`),
		g.Group(children),
	)
}