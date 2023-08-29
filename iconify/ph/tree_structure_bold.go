package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeStructureBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 116h48a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20h-48a20 20 0 0 0-20 20v12h-4a36 36 0 0 0-36 36v20H84v-4a20 20 0 0 0-20-20H32a20 20 0 0 0-20 20v32a20 20 0 0 0 20 20h32a20 20 0 0 0 20-20v-4h24v20a36 36 0 0 0 36 36h4v12a20 20 0 0 0 20 20h48a20 20 0 0 0 20-20v-48a20 20 0 0 0-20-20h-48a20 20 0 0 0-20 20v12h-4a12 12 0 0 1-12-12V96a12 12 0 0 1 12-12h4v12a20 20 0 0 0 20 20ZM60 140H36v-24h24Zm112 24h40v40h-40Zm0-112h40v40h-40Z"/>`),
		g.Group(children),
	)
}