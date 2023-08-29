package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeStructureThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 108h48a12 12 0 0 0 12-12V48a12 12 0 0 0-12-12h-48a12 12 0 0 0-12 12v20h-12a28 28 0 0 0-28 28v28H76v-12a12 12 0 0 0-12-12H32a12 12 0 0 0-12 12v32a12 12 0 0 0 12 12h32a12 12 0 0 0 12-12v-12h40v28a28 28 0 0 0 28 28h12v20a12 12 0 0 0 12 12h48a12 12 0 0 0 12-12v-48a12 12 0 0 0-12-12h-48a12 12 0 0 0-12 12v20h-12a20 20 0 0 1-20-20V96a20 20 0 0 1 20-20h12v20a12 12 0 0 0 12 12ZM68 144a4 4 0 0 1-4 4H32a4 4 0 0 1-4-4v-32a4 4 0 0 1 4-4h32a4 4 0 0 1 4 4Zm96 16a4 4 0 0 1 4-4h48a4 4 0 0 1 4 4v48a4 4 0 0 1-4 4h-48a4 4 0 0 1-4-4Zm0-112a4 4 0 0 1 4-4h48a4 4 0 0 1 4 4v48a4 4 0 0 1-4 4h-48a4 4 0 0 1-4-4Z"/>`),
		g.Group(children),
	)
}