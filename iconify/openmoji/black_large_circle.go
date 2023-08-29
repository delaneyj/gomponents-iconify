package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackLargeCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path stroke="#000" stroke-linejoin="round" stroke-width="2" d="M36 66c16.569 0 30-13.431 30-30C66 19.432 52.569 6 36 6C19.431 6 6 19.432 6 36c0 16.57 13.431 30 30 30Z"/><path fill="#3F3F3F" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M36 66c16.569 0 30-13.431 30-30C66 19.432 52.569 6 36 6C19.431 6 6 19.432 6 36c0 16.57 13.431 30 30 30Z"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M36 66c16.569 0 30-13.431 30-30C66 19.432 52.569 6 36 6C19.431 6 6 19.432 6 36c0 16.57 13.431 30 30 30Z"/>`),
		g.Group(children),
	)
}