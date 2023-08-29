package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCloudThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m210.83 85.17l-56-56A4 4 0 0 0 152 28H56a12 12 0 0 0-12 12v88a4 4 0 0 0 8 0V40a4 4 0 0 1 4-4h92v52a4 4 0 0 0 4 4h52v124a4 4 0 0 1-4 4h-24a4 4 0 0 0 0 8h24a12 12 0 0 0 12-12V88a4 4 0 0 0-1.17-2.83ZM156 41.65L198.34 84H156ZM108 132a47.72 47.72 0 0 0-45.3 32H60a32 32 0 0 0 0 64h48a48 48 0 0 0 0-96Zm0 88H60a24 24 0 0 1 0-48h.66c-.2 1.2-.35 2.41-.46 3.64a4 4 0 0 0 8 .72a41.2 41.2 0 0 1 1.23-6.92a4.68 4.68 0 0 0 .21-.73A40 40 0 1 1 108 220Z"/>`),
		g.Group(children),
	)
}