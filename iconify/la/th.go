package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Th(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 6v20h26V6zm2 2h4v4H5zm6 0h4v4h-4zm6 0h4v4h-4zm6 0h4v4h-4zM5 14h4v4H5zm6 0h4v4h-4zm6 0h4v4h-4zm6 0h4v4h-4zM5 20h4v4H5zm6 0h4v4h-4zm6 0h4v4h-4zm6 0h4v4h-4z"/>`),
		g.Group(children),
	)
}