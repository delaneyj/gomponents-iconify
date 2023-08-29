package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VinylRecordThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm0-152a60.07 60.07 0 0 0-60 60a4 4 0 0 1-8 0a68.07 68.07 0 0 1 68-68a4 4 0 0 1 0 8Zm68 60a68.07 68.07 0 0 1-68 68a4 4 0 0 1 0-8a60.07 60.07 0 0 0 60-60a4 4 0 0 1 8 0Zm-40 0a28 28 0 1 0-28 28a28 28 0 0 0 28-28Zm-48 0a20 20 0 1 1 20 20a20 20 0 0 1-20-20Z"/>`),
		g.Group(children),
	)
}