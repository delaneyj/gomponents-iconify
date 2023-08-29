package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudWarningThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 44a84.11 84.11 0 0 0-76.41 49.12A60.71 60.71 0 0 0 72 92a60 60 0 0 0 0 120h88a84 84 0 0 0 0-168Zm0 160H72a52 52 0 1 1 8.55-103.3A83.66 83.66 0 0 0 76 128a4 4 0 0 0 8 0a76 76 0 1 1 76 76Zm-4-76V88a4 4 0 0 1 8 0v40a4 4 0 0 1-8 0Zm12 36a8 8 0 1 1-8-8a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}