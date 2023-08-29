package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TentLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.103 19L12.866 3a1 1 0 0 0-1.732 0L1.896 19H1v2h22v-2h-.897ZM7.6 19H4.206L12 5.5L19.794 19H16.4L12 11l-4.4 8Zm4.4-3.85L14.117 19H9.883L12 15.15Z"/>`),
		g.Group(children),
	)
}