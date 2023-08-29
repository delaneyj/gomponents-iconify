package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePptThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 152a4 4 0 0 1-4 4h-16v52a4 4 0 0 1-8 0v-52h-16a4 4 0 0 1 0-8h40a4 4 0 0 1 4 4ZM88 172a24 24 0 0 1-24 24H52v12a4 4 0 0 1-8 0v-56a4 4 0 0 1 4-4h16a24 24 0 0 1 24 24Zm-8 0a16 16 0 0 0-16-16H52v32h12a16 16 0 0 0 16-16Zm76 0a24 24 0 0 1-24 24h-12v12a4 4 0 0 1-8 0v-56a4 4 0 0 1 4-4h16a24 24 0 0 1 24 24Zm-8 0a16 16 0 0 0-16-16h-12v32h12a16 16 0 0 0 16-16ZM44 112V40a12 12 0 0 1 12-12h96a4 4 0 0 1 2.83 1.17l56 56A4 4 0 0 1 212 88v24a4 4 0 0 1-8 0V92h-52a4 4 0 0 1-4-4V36H56a4 4 0 0 0-4 4v72a4 4 0 0 1-8 0Zm112-28h42.34L156 41.65Z"/>`),
		g.Group(children),
	)
}