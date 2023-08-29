package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.5 11H0v4h15v-4H7.5zm6.5 2h-2v-1h2v1zM3.5 5l4-4l4 4H9v5H6V5z"/>`),
		g.Group(children),
	)
}