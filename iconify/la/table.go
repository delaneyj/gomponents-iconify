package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h5v5H7zm7 0h4v5h-4zm6 0h5v5h-5zM7 14h5v4H7zm7 0h4v4h-4zm6 0h5v4h-5zM7 20h5v5H7zm7 0h4v5h-4zm6 0h5v5h-5z"/>`),
		g.Group(children),
	)
}