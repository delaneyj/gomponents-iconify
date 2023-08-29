package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoctorTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.75 4.498a.25.25 0 0 0-.25.25V9.75a.75.75 0 0 1-.75.75h-5a.25.25 0 0 0-.25.25v6.5c0 .138.112.25.25.25h5a.75.75 0 0 1 .75.75v5.001c0 .138.112.25.25.25h6.5a.25.25 0 0 0 .25-.25V18.25a.75.75 0 0 1 .75-.75h5a.25.25 0 0 0 .25-.25v-6.5a.25.25 0 0 0-.25-.25h-5a.75.75 0 0 1-.75-.75V4.749a.25.25 0 0 0-.25-.25h-6.5ZM9 4.748c0-.966.784-1.75 1.75-1.75h6.5c.966 0 1.75.784 1.75 1.75V9h4.25c.966 0 1.75.783 1.75 1.75v6.5A1.75 1.75 0 0 1 23.25 19H19v4.251a1.75 1.75 0 0 1-1.75 1.75h-6.5A1.75 1.75 0 0 1 9 23.251V19H4.75A1.75 1.75 0 0 1 3 17.25v-6.5C3 9.784 3.784 9 4.75 9H9V4.749Z"/>`),
		g.Group(children),
	)
}