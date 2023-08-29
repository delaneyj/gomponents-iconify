package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeTransparent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m221.66 90.34l-56-56A8 8 0 0 0 160 32H40a8 8 0 0 0-8 8v120a8 8 0 0 0 2.3 5.61l56 56A8 8 0 0 0 96 224h120a8 8 0 0 0 8-8V96a8 8 0 0 0-2.34-5.66ZM168 59.31L196.69 88H168ZM88 196.69L59.31 168H88ZM88 152H48V59.31l40 40ZM59.31 48H152v40H99.31ZM152 104v48h-48v-48Zm-48 104v-40h52.69l40 40Zm104-11.31l-40-40V104h40Z"/>`),
		g.Group(children),
	)
}