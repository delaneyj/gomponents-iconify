package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControllerVolume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 13.805c0 .657-.538 1.195-1.195 1.195H1.533c-.88 0-.982-.371-.229-.822l16.323-9.055C18.382 4.67 19 5.019 19 5.9v7.905z"/>`),
		g.Group(children),
	)
}