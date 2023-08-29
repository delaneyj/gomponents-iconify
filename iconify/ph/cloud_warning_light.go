package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudWarningLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 42a86.11 86.11 0 0 0-77.57 48.88A62 62 0 1 0 72 214h88a86 86 0 0 0 0-172Zm0 160H72a50 50 0 0 1 0-100a50.67 50.67 0 0 1 5.91.35A85.61 85.61 0 0 0 74 128a6 6 0 0 0 12 0a74 74 0 1 1 74 74Zm-6-74V88a6 6 0 0 1 12 0v40a6 6 0 0 1-12 0Zm16 36a10 10 0 1 1-10-10a10 10 0 0 1 10 10Z"/>`),
		g.Group(children),
	)
}