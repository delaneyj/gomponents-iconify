package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cameraflash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m640 768l-352 256l-160-256h149l64-192H0L192 0h192L235 448h341L469 768h171z"/>`),
		g.Group(children),
	)
}