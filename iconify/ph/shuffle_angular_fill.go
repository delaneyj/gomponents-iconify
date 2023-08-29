package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleAngularFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M237.66 178.34a8 8 0 0 1 0 11.32l-24 24A8 8 0 0 1 200 208v-16h-27.88a16 16 0 0 1-13-6.7L83.88 80H32a8 8 0 0 1 0-16h51.88a16 16 0 0 1 13 6.7L172.12 176H200v-16a8 8 0 0 1 13.66-5.66ZM143 107a8 8 0 0 0 11.16-1.86l18-25.12H200V96a8 8 0 0 0 13.66 5.66l24-24a8 8 0 0 0 0-11.32l-24-24A8 8 0 0 0 200 48v16h-27.88a16 16 0 0 0-13 6.7l-17.97 25.12A8 8 0 0 0 143 107Zm-30 42a8 8 0 0 0-11.16 1.86L83.88 176H32a8 8 0 0 0 0 16h51.88a16 16 0 0 0 13-6.7l17.95-25.12A8 8 0 0 0 113 149Z"/>`),
		g.Group(children),
	)
}