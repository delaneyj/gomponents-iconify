package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelTopGalleryTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.754 4a3.75 3.75 0 0 0-3.75 3.75v12.5A3.75 3.75 0 0 0 5.754 24H22.25A3.75 3.75 0 0 0 26 20.25V7.75A3.75 3.75 0 0 0 22.25 4H5.755Zm-2.25 3.75a2.25 2.25 0 0 1 2.25-2.25H9.5v6H3.504V7.75Zm0 5.25H24.5v7.25a2.25 2.25 0 0 1-2.25 2.25H5.755a2.25 2.25 0 0 1-2.25-2.25V13ZM24.5 11.5h-6v-6h3.75a2.25 2.25 0 0 1 2.25 2.25v3.75Zm-7.5-6v6h-6v-6h6Z"/>`),
		g.Group(children),
	)
}