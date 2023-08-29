package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatboxEllipsesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M456 48H56a24 24 0 0 0-24 24v288a24 24 0 0 0 24 24h72v80l117.74-80H456a24 24 0 0 0 24-24V72a24 24 0 0 0-24-24ZM160 248a32 32 0 1 1 32-32a32 32 0 0 1-32 32Zm96 0a32 32 0 1 1 32-32a32 32 0 0 1-32 32Zm96 0a32 32 0 1 1 32-32a32 32 0 0 1-32 32ZM456 80Z"/>`),
		g.Group(children),
	)
}