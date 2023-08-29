package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeMute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.885.133L7.094 5.188v13.625l8.791 5.055c.726.416 1.572-.202 1.572-1.145V1.279c0-.945-.847-1.564-1.572-1.145zM6.003 5.727l-4.926.01A1.184 1.184 0 0 0 .004 7v-.004v10.256a1.167 1.167 0 0 0 1.059 1.246h4.939V5.726z"/>`),
		g.Group(children),
	)
}