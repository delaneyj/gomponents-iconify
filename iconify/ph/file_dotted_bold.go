package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDottedBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 92v-4a12.3 12.3 0 0 0-3.5-8.5l-56-56A12.3 12.3 0 0 0 152 20h-32a12 12 0 0 0 0 24h16v48a12 12 0 0 0 12 12h48v32a12 12 0 0 0 24 0V92Zm-60-35l23 23h-23ZM84 224a12 12 0 0 1-12 12H56a20.1 20.1 0 0 1-20-20v-32a12 12 0 0 1 24 0v28h12a12 12 0 0 1 12 12Zm8-192a12 12 0 0 1-12 12H60v20a12 12 0 0 1-24 0V40a20.1 20.1 0 0 1 20-20h24a12 12 0 0 1 12 12Zm128 144v40a20.1 20.1 0 0 1-20 20h-8a12 12 0 0 1 0-24h4v-36a12 12 0 0 1 24 0ZM36 144v-40a12 12 0 0 1 24 0v40a12 12 0 0 1-24 0Zm128 80a12 12 0 0 1-12 12h-40a12 12 0 0 1 0-24h40a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}