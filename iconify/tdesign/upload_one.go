package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2.586L18.414 9L17 10.414l-4-4V16h-2V6.414l-4 4L5.586 9L12 2.586ZM3 18h18v2H3v-2Z"/>`),
		g.Group(children),
	)
}