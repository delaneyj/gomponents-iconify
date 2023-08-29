package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeTransparentFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M104 152v-48h48v48ZM32 53v95a4 4 0 0 0 4 4h52V99.31L38.83 50.14A4 4 0 0 0 32 53Zm188 51h-52v52.69l49.17 49.17A4 4 0 0 0 224 203v-95a4 4 0 0 0-4-4Zm-68-68a4 4 0 0 0-4-4H53a4 4 0 0 0-2.83 6.83L99.31 88H152Zm60.49 45.17l-37.66-37.66a4 4 0 0 0-6.83 2.83V88h41.66a4 4 0 0 0 2.83-6.83ZM156.69 168H104v52a4 4 0 0 0 4 4h95a4 4 0 0 0 2.83-6.83Zm-113.18 6.83l37.66 37.66a4 4 0 0 0 6.83-2.83V168H46.34a4 4 0 0 0-2.83 6.83Z"/>`),
		g.Group(children),
	)
}