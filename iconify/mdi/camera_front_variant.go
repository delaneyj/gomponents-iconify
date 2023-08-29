package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraFrontVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 0h12a2 2 0 0 1 2 2v20a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2m6 6a3 3 0 0 1 3 3a3 3 0 0 1-3 3a3 3 0 0 1-3-3a3 3 0 0 1 3-3m-1-5v2h2V1h-2M6 4v12.5c0-1.38 2.69-2.5 6-2.5s6 1.12 6 2.5V4H6m7 14H9v2h4v2l3-3l-3-3v2Z"/>`),
		g.Group(children),
	)
}