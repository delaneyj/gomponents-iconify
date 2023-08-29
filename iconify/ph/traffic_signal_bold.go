package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficSignalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 140h-12V84h12a12 12 0 0 0 0-24h-12V40a20 20 0 0 0-20-20H72a20 20 0 0 0-20 20v20H40a12 12 0 0 0 0 24h12v56H40a12 12 0 0 0 0 24h12v52a20 20 0 0 0 20 20h112a20 20 0 0 0 20-20v-52h12a12 12 0 0 0 0-24Zm-36 72H76V44h104Zm-52-92a32 32 0 1 0-32-32a32 32 0 0 0 32 32Zm0-40a8 8 0 1 1-8 8a8 8 0 0 1 8-8Zm0 120a32 32 0 1 0-32-32a32 32 0 0 0 32 32Zm0-40a8 8 0 1 1-8 8a8 8 0 0 1 8-8Z"/>`),
		g.Group(children),
	)
}