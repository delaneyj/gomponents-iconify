package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M36 64c15.464 0 28-12.536 28-28S51.464 8 36 8S8 20.536 8 36s12.536 28 28 28Z"/><path fill="#3F3F3F" d="M36 64c15.464 0 28-12.536 28-28S51.464 8 36 8S8 20.536 8 36s12.536 28 28 28Z"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M36 64c15.464 0 28-12.536 28-28S51.464 8 36 8S8 20.536 8 36s12.536 28 28 28Z"/>`),
		g.Group(children),
	)
}