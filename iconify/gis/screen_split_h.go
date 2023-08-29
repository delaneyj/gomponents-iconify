package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenSplitH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M2.5 12.5A2.5 2.5 0 0 0 0 15v70a2.5 2.5 0 0 0 2.5 2.5h95A2.5 2.5 0 0 0 100 85V15a2.5 2.5 0 0 0-2.5-2.5zm2.5 5h90v65H5Z" color="currentColor"/><path fill="currentColor" d="M43 21v23.873h-8.604V35L20 50l14.396 15v-9.877H43V79h4.5V21H43zm9.5 0v58H57V55.125h8.605v9.873L80 50L65.605 35.004v9.869H57V21h-4.5z" color="currentColor"/>`),
		g.Group(children),
	)
}