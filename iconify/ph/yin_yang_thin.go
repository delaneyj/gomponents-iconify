package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYangThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28ZM36 128a92.1 92.1 0 0 1 92-92a44 44 0 0 1 0 88a52 52 0 0 0-37.44 88A92.14 92.14 0 0 1 36 128Zm92 92a44 44 0 0 1 0-88a52 52 0 0 0 37.44-88A92 92 0 0 1 128 220Zm8-44a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm-16-96a8 8 0 1 1 8 8a8 8 0 0 1-8-8Z"/>`),
		g.Group(children),
	)
}