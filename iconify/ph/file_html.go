package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileHtml(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 152a8 8 0 0 1-8 8h-8v48a8 8 0 0 1-16 0v-48h-8a8 8 0 0 1 0-16h32a8 8 0 0 1 8 8Zm-64-8a8 8 0 0 0-8 8v20H40v-20a8 8 0 0 0-16 0v56a8 8 0 0 0 16 0v-20h16v20a8 8 0 0 0 16 0v-56a8 8 0 0 0-8-8Zm176 56h-12v-48a8 8 0 0 0-16 0v56a8 8 0 0 0 8 8h20a8 8 0 0 0 0-16Zm-45.86-55.71a8 8 0 0 0-9 3.59L168 176.45l-17.14-28.57A8 8 0 0 0 136 152v56a8 8 0 0 0 16 0v-27.12l9.14 15.24a8 8 0 0 0 13.72 0l9.14-15.24V208a8 8 0 0 0 16 0v-56a8 8 0 0 0-5.86-7.71ZM208 120a8 8 0 0 1-8-8V96h-48a8 8 0 0 1-8-8V40H56v72a8 8 0 0 1-16 0V40a16 16 0 0 1 16-16h96a8 8 0 0 1 5.66 2.34l56 56A8 8 0 0 1 216 88v24a8 8 0 0 1-8 8Zm-19.31-40L160 51.31V80Z"/>`),
		g.Group(children),
	)
}