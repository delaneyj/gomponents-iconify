package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpenTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 16.137A2.742 2.742 0 0 1 8 17H3.75A1.75 1.75 0 0 1 2 15.25V4.75C2 3.784 2.784 3 3.75 3H8c.788 0 1.499.331 2 .863A2.742 2.742 0 0 1 12 3h4.25c.966 0 1.75.784 1.75 1.75v10.5A1.75 1.75 0 0 1 16.25 17H12a2.742 2.742 0 0 1-2-.863ZM3.5 4.75v10.5c0 .138.112.25.25.25H8c.69 0 1.25-.56 1.25-1.25v-8.5c0-.69-.56-1.25-1.25-1.25H3.75a.25.25 0 0 0-.25.25Zm7.25 9.5c0 .69.56 1.25 1.25 1.25h4.25a.25.25 0 0 0 .25-.25V4.75a.25.25 0 0 0-.25-.25H12c-.69 0-1.25.56-1.25 1.25v8.5Z"/>`),
		g.Group(children),
	)
}