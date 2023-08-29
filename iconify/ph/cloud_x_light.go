package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudXLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 42a86.11 86.11 0 0 0-77.57 48.88A62 62 0 1 0 72 214h88a86 86 0 0 0 0-172Zm0 160H72a50 50 0 0 1 0-100a50.67 50.67 0 0 1 5.91.35A85.61 85.61 0 0 0 74 128a6 6 0 0 0 12 0a74 74 0 1 1 74 74Zm28.24-85.76L168.48 136l19.76 19.76a6 6 0 1 1-8.48 8.48L160 144.48l-19.76 19.76a6 6 0 0 1-8.48-8.48L151.52 136l-19.76-19.76a6 6 0 0 1 8.48-8.48L160 127.52l19.76-19.76a6 6 0 0 1 8.48 8.48Z"/>`),
		g.Group(children),
	)
}