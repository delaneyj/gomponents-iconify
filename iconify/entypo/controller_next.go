package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControllerNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.244 9.52L5.041 4.571C4.469 4.188 4 4.469 4 5.196v9.609c0 .725.469 1.006 1.041.625l7.203-4.951s.279-.199.279-.478c0-.28-.279-.481-.279-.481zM14 4h1c.553 0 1 .048 1 .6v10.8c0 .552-.447.6-1 .6h-1c-.553 0-1-.048-1-.6V4.6c0-.552.447-.6 1-.6z"/>`),
		g.Group(children),
	)
}