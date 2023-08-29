package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDesktopMobileTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.75 2A2.25 2.25 0 0 1 18 4.25v15.5A2.25 2.25 0 0 1 15.75 22h-7.5A2.25 2.25 0 0 1 6 19.75V4.25A2.25 2.25 0 0 1 8.25 2h7.5Zm-2.5 16h-2.5a.75.75 0 0 0-.102 1.493l.102.007h2.5a.75.75 0 0 0 0-1.5Zm1.25-9.001h-5a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h5a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5ZM14.5 5h-5a.5.5 0 0 0-.5.5v1.997a.5.5 0 0 0 .5.5h5a.5.5 0 0 0 .5-.5V5.5a.5.5 0 0 0-.5-.5Z"/>`),
		g.Group(children),
	)
}