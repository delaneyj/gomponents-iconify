package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsersBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M125.18 156.94a64 64 0 1 0-82.36 0a100.23 100.23 0 0 0-39.49 32a12 12 0 0 0 19.35 14.2a76 76 0 0 1 122.64 0a12 12 0 0 0 19.36-14.2a100.33 100.33 0 0 0-39.5-32ZM44 108a40 40 0 1 1 40 40a40 40 0 0 1-40-40Zm206.1 97.67a12 12 0 0 1-16.78-2.57A76.31 76.31 0 0 0 172 172a12 12 0 0 1 0-24a40 40 0 1 0-14.85-77.16a12 12 0 1 1-8.92-22.28a64 64 0 0 1 65 108.38a100.23 100.23 0 0 1 39.49 32a12 12 0 0 1-2.62 16.73Z"/>`),
		g.Group(children),
	)
}