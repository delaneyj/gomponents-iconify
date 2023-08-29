package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorefrontFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 96a7.89 7.89 0 0 0-.3-2.2l-14.35-50.2A16.07 16.07 0 0 0 202 32H54a16.07 16.07 0 0 0-15.35 11.6L24.31 93.8A7.89 7.89 0 0 0 24 96v16a40 40 0 0 0 16 32v64a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16v-64a40 40 0 0 0 16-32V96ZM88 112a24 24 0 0 1-35.12 21.26a7.88 7.88 0 0 0-1.82-1.06A24 24 0 0 1 40 112v-8h48Zm64 0a24 24 0 0 1-48 0v-8h48Zm64 0a24 24 0 0 1-11.07 20.2a8.08 8.08 0 0 0-1.8 1.05A24 24 0 0 1 168 112v-8h48Z"/>`),
		g.Group(children),
	)
}