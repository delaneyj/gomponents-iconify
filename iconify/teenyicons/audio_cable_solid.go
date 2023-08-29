package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioCableSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.854.146a.5.5 0 0 0-.708 0l-1 1A.5.5 0 0 0 6 1.5V3h3V1.5a.5.5 0 0 0-.146-.354l-1-1ZM9 4v4H6V4h3Zm1 5v3.5A1.5 1.5 0 0 1 8.5 14H8v1H7v-1h-.5A1.5 1.5 0 0 1 5 12.5V9h5Z"/>`),
		g.Group(children),
	)
}