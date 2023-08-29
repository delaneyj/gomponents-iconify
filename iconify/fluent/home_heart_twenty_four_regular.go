package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeHeartTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.55 2.532a2.25 2.25 0 0 1 2.9 0l6.75 5.692c.507.428.8 1.057.8 1.72v1.317A4.397 4.397 0 0 0 19.5 11V9.944a.75.75 0 0 0-.267-.573l-6.75-5.692a.75.75 0 0 0-.966 0L4.767 9.37a.75.75 0 0 0-.267.573v9.31c0 .138.112.25.25.25h7.169l1.449 1.5H4.75A1.75 1.75 0 0 1 3 19.254v-9.31c0-.663.293-1.292.8-1.72l6.75-5.692Zm11.427 15.64c1.364-1.411 1.364-3.7 0-5.113a3.547 3.547 0 0 0-.476-.413a3.406 3.406 0 0 0-4.465.413h-.072A3.408 3.408 0 0 0 13 12.346a3.566 3.566 0 0 0-1.5 1.404c-.81 1.39-.636 3.223.523 4.423l4.442 4.598a.738.738 0 0 0 1.07 0l4.442-4.598Z"/>`),
		g.Group(children),
	)
}