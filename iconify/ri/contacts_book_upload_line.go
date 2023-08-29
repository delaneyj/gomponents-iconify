package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactsBookUploadLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.005 2C20.107 2 21 2.898 21 3.99v16.02c0 1.099-.893 1.99-1.995 1.99H3V2h16.005ZM7 4H5v16h2V4Zm12 0H9v16h10V4Zm-5 4l4 4h-3v4h-2v-4h-3l4-4Zm10 4v4h-2v-4h2Zm0-6v4h-2V6h2Z"/>`),
		g.Group(children),
	)
}