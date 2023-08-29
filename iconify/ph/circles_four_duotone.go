package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesFourDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M112 80a32 32 0 1 1-32-32a32 32 0 0 1 32 32Zm64 32a32 32 0 1 0-32-32a32 32 0 0 0 32 32Zm-96 32a32 32 0 1 0 32 32a32 32 0 0 0-32-32Zm96 0a32 32 0 1 0 32 32a32 32 0 0 0-32-32Z" opacity=".2"/><path d="M80 40a40 40 0 1 0 40 40a40 40 0 0 0-40-40Zm0 64a24 24 0 1 1 24-24a24 24 0 0 1-24 24Zm96 16a40 40 0 1 0-40-40a40 40 0 0 0 40 40Zm0-64a24 24 0 1 1-24 24a24 24 0 0 1 24-24Zm-96 80a40 40 0 1 0 40 40a40 40 0 0 0-40-40Zm0 64a24 24 0 1 1 24-24a24 24 0 0 1-24 24Zm96-64a40 40 0 1 0 40 40a40 40 0 0 0-40-40Zm0 64a24 24 0 1 1 24-24a24 24 0 0 1-24 24Z"/></g>`),
		g.Group(children),
	)
}