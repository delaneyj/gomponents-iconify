package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 50H32a6 6 0 0 0-6 6v144a6 6 0 0 0 6 6h192a6 6 0 0 0 6-6V56a6 6 0 0 0-6-6ZM86 146v-36h84v36Zm-48 0v-36h36v36Zm144-36h36v36h-36Zm36-12h-84V62h84Zm-96-36v36H38V62Zm-84 96h84v36H38Zm96 36v-36h84v36Z"/>`),
		g.Group(children),
	)
}