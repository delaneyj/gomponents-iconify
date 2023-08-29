package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelRightGalleryTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22 7.25A3.25 3.25 0 0 0 18.75 4H5.25A3.25 3.25 0 0 0 2 7.25v9.5A3.25 3.25 0 0 0 5.25 20h13.5A3.25 3.25 0 0 0 22 16.75v-9.5Zm-1.5 6.25h-5v-3h5v3Zm-5 1.5h5v1.75a1.75 1.75 0 0 1-1.75 1.75H15.5V15Zm5-7.75V9h-5V5.5h3.25c.966 0 1.75.784 1.75 1.75ZM14 5.5v13H5.25a1.75 1.75 0 0 1-1.75-1.75v-9.5c0-.966.784-1.75 1.75-1.75H14Z"/>`),
		g.Group(children),
	)
}