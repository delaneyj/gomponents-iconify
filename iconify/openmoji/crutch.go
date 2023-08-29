package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crutch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="#61b2e4"><path d="M37.809 65.826c0-2.387-1.079-4.322-2.41-4.322s-2.409 1.935-2.409 4.322h4.819ZM46.31 26.29a.999.999 0 0 1-1.27-.63l-.27-.81l-6.88 2.38l-1.89.66l-.65-1.89l-5.57-16.11c-.18-.52.1-1.09.62-1.27c.52-.18 1.09.1 1.27.62l5.57 16.1l7.83-2.71c.25-.09.53-.07.77.05c.24.12.42.32.51.58l.59 1.77c.18.52-.1 1.09-.63 1.26Z"/><path d="m28.448 8.932l7.399-2.55a3.787 3.787 0 0 1 4.812 2.344a3.787 3.787 0 0 1-2.345 4.812l-7.4 2.55l-2.466-7.156Z"/></g><g fill="none" stroke="#000" stroke-width="2"><path stroke-miterlimit="10" d="M35.399 29.098v32.406"/><path stroke-linecap="round" stroke-linejoin="round" d="M37.809 65.826c0-2.387-1.079-4.322-2.41-4.322s-2.409 1.935-2.409 4.322h4.819ZM28.448 8.932l7.399-2.55a3.787 3.787 0 0 1 4.812 2.344h0a3.787 3.787 0 0 1-2.345 4.812l-7.4 2.55l-2.466-7.156h0Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M45.993 25.44L35.41 29.098L28.448 8.954"/></g>`),
		g.Group(children),
	)
}