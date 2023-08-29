package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadCloudFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20.981a6.5 6.5 0 0 1-2.936-12a8.001 8.001 0 0 1 15.872 0a6.5 6.5 0 0 1-2.936 12V21H7v-.019ZM13 12V8h-2v4H8l4 5l4-5h-3Z"/>`),
		g.Group(children),
	)
}