package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m11.5 7l-4 4l-4-4H6V1h3v6zm-4 4H0v4h15v-4H7.5zm6.5 2h-2v-1h2v1z"/>`),
		g.Group(children),
	)
}