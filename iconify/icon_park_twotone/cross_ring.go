package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossRing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCrossRing0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 32a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM8 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm32 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="M12 24h24M24 12v24m13.305-20.89a16.073 16.073 0 0 0-4.415-4.415m-17.78 0a16.074 16.074 0 0 0-4.415 4.415l4.415-4.415ZM10.696 32.89a16.08 16.08 0 0 0 4.415 4.415l-4.415-4.415Zm22.195 4.415a16.079 16.079 0 0 0 4.415-4.415l-4.415 4.415Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCrossRing0)"/>`),
		g.Group(children),
	)
}