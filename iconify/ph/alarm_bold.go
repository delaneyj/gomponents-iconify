package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 176a76 76 0 1 1 76-76a76.08 76.08 0 0 1-76 76ZM32.49 64.49a12 12 0 1 1-17-17l32-32a12 12 0 1 1 17 17Zm208 0a12 12 0 0 1-17 0l-32-32a12 12 0 0 1 17-17l32 32a12 12 0 0 1 0 17ZM176 116a12 12 0 0 1 0 24h-48a12 12 0 0 1-12-12V80a12 12 0 0 1 24 0v36Z"/>`),
		g.Group(children),
	)
}