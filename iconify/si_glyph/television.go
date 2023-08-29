package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Television(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.381 5.06H9.468l2.739-3.79l-1.012-.223L8.303 5.06h-.008L5.403 1.047l-1.014.223l2.739 3.79H.656a.582.582 0 0 0-.593.571v10.785c0 .317.265.572.593.572h14.723a.582.582 0 0 0 .594-.572V5.631a.578.578 0 0 0-.592-.571zM13 14.67c0 .206-.152.372-.342.372H2.263c-.188 0-.341-.166-.341-.372V7.304c0-.204.152-.372.341-.372h10.395c.189 0 .342.168.342.372v7.366zm2 .372h-1V14h1v1.042z"/>`),
		g.Group(children),
	)
}