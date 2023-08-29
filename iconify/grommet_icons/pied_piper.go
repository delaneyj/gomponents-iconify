package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PiedPiper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 19.421c2.274 0 4.042-.758 4.042-.758s3.032-7.579 7.326-7.579c3.285 0 3.79 2.527 3.79 2.527S19.958 4.263 24 3c-3.79 3.032-3.284 6.316-5.053 7.832c-1.768 1.515-1.768.006-3.79 3.543c-4.546.505-6.032 2.014-9.094 3.783c5.305-2.526 6.316-2.78 11.116-2.526c.504.026.758.252.505.757c-.733 1.466-1.28 3.673-2.273 3.537c-5.558-.758-8.843.506-11.622.506c-2.778 0-3.789-.506-3.789-1.01Z"/>`),
		g.Group(children),
	)
}