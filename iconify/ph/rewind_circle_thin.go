package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm45.89-127.53a4 4 0 0 0-4.11.2l-48 32a4 4 0 0 0 0 6.66l48 32A4 4 0 0 0 176 160V96a4 4 0 0 0-2.11-3.53ZM168 152.53L131.21 128L168 103.47Zm-50.11-60.06a4 4 0 0 0-4.11.2l-48 32a4 4 0 0 0 0 6.66l48 32A4 4 0 0 0 120 160V96a4 4 0 0 0-2.11-3.53ZM112 152.53L75.21 128L112 103.47Z"/>`),
		g.Group(children),
	)
}