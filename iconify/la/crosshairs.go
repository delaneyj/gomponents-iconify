package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crosshairs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 3v2.063C9.734 5.538 5.54 9.733 5.062 15H3v2h2.063c.476 5.266 4.671 9.46 9.937 9.938V29h2v-2.063c5.266-.476 9.46-4.671 9.938-9.937H29v-2h-2.063c-.475-5.266-4.67-9.46-9.937-9.938V3zm0 4.031V9h2V7.031A8.999 8.999 0 0 1 24.969 15H23v2h1.969A8.999 8.999 0 0 1 17 24.969V23h-2v1.969A8.999 8.999 0 0 1 7.031 17H9v-2H7.031A8.999 8.999 0 0 1 15 7.031z"/>`),
		g.Group(children),
	)
}