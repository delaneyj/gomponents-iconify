package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectionAllBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M100 40a12 12 0 0 1 12-12h32a12 12 0 0 1 0 24h-32a12 12 0 0 1-12-12Zm44 164h-32a12 12 0 0 0 0 24h32a12 12 0 0 0 0-24Zm64-176h-24a12 12 0 0 0 0 24h20v20a12 12 0 0 0 24 0V48a20 20 0 0 0-20-20Zm8 72a12 12 0 0 0-12 12v32a12 12 0 0 0 24 0v-32a12 12 0 0 0-12-12Zm0 72a12 12 0 0 0-12 12v20h-20a12 12 0 0 0 0 24h24a20 20 0 0 0 20-20v-24a12 12 0 0 0-12-12ZM40 156a12 12 0 0 0 12-12v-32a12 12 0 0 0-24 0v32a12 12 0 0 0 12 12Zm32 48H52v-20a12 12 0 0 0-24 0v24a20 20 0 0 0 20 20h24a12 12 0 0 0 0-24ZM40 84a12 12 0 0 0 12-12V52h20a12 12 0 0 0 0-24H48a20 20 0 0 0-20 20v24a12 12 0 0 0 12 12Zm40-16h96a12 12 0 0 1 12 12v96a12 12 0 0 1-12 12H80a12 12 0 0 1-12-12V80a12 12 0 0 1 12-12Zm12 96h72V92H92Z"/>`),
		g.Group(children),
	)
}