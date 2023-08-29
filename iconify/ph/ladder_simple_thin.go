package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LadderSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 28a4 4 0 0 0-4 4v36H68V32a4 4 0 0 0-8 0v192a4 4 0 0 0 8 0v-36h120v36a4 4 0 0 0 8 0V32a4 4 0 0 0-4-4Zm-4 48v48H68V76ZM68 180v-48h120v48Z"/>`),
		g.Group(children),
	)
}