package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unacademy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.916 7.534h38.168m-38.584 0a19.5 19.5 0 0 0 39 0Zm28.914 32.932a9.414 9.414 0 0 0-18.828 0Z"/>`),
		g.Group(children),
	)
}