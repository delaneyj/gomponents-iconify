package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdBarcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M88 128h48v256H88z" fill="currentColor"/><path d="M232 128h48v256h-48z" fill="currentColor"/><path d="M160 144h48v224h-48z" fill="currentColor"/><path d="M304 144h48v224h-48z" fill="currentColor"/><path d="M376 128h48v256h-48z" fill="currentColor"/><path d="M104 104V56H16v400h88v-48H64V104z" fill="currentColor"/><path d="M408 56v48h40v304h-40v48h88V56z" fill="currentColor"/>`),
		g.Group(children),
	)
}