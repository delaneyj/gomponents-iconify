package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelTopGalleryTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.25 4A3.25 3.25 0 0 0 2 7.25v9.5A3.25 3.25 0 0 0 5.25 20h13.5A3.25 3.25 0 0 0 22 16.75v-9.5A3.25 3.25 0 0 0 18.75 4H5.25ZM3.5 7.25c0-.966.784-1.75 1.75-1.75H8v5H3.5V7.25Zm0 4.75h17v4.75a1.75 1.75 0 0 1-1.75 1.75H5.25a1.75 1.75 0 0 1-1.75-1.75V12Zm17-1.5H16v-5h2.75c.966 0 1.75.784 1.75 1.75v3.25Zm-6-5v5h-5v-5h5Z"/>`),
		g.Group(children),
	)
}