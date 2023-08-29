package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m247.35 148.85l-48-112A8 8 0 0 0 192 32H64a8 8 0 0 0-7.35 4.85l-48 112A8 8 0 0 0 16 160h104v48H96a8 8 0 0 0 0 16h64a8 8 0 0 0 0-16h-24v-48h56v32a8 8 0 0 0 16 0v-32h32a8 8 0 0 0 7.35-11.15ZM28.13 144l41.15-96h117.44l41.15 96Z"/>`),
		g.Group(children),
	)
}