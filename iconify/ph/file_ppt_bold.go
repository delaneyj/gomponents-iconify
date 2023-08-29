package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePptBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 152a12 12 0 0 1-12 12h-8v44a12 12 0 0 1-24 0v-44h-8a12 12 0 0 1 0-24h40a12 12 0 0 1 12 12ZM92 172a32 32 0 0 1-32 32h-4v4a12 12 0 0 1-24 0v-56a12 12 0 0 1 12-12h16a32 32 0 0 1 32 32Zm-24 0a8 8 0 0 0-8-8h-4v16h4a8 8 0 0 0 8-8Zm96 0a32 32 0 0 1-32 32h-4v4a12 12 0 0 1-24 0v-56a12 12 0 0 1 12-12h16a32 32 0 0 1 32 32Zm-24 0a8 8 0 0 0-8-8h-4v16h4a8 8 0 0 0 8-8ZM36 108V40a20 20 0 0 1 20-20h96a12 12 0 0 1 8.49 3.52l56 56A12 12 0 0 1 220 88v20a12 12 0 0 1-24 0v-4h-48a12 12 0 0 1-12-12V44H60v64a12 12 0 0 1-24 0Zm124-28h23l-23-23Z"/>`),
		g.Group(children),
	)
}