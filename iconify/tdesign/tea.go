package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2v5H9V2h2ZM7 3v4H5V3h2Zm8 0v4h-2V3h2ZM2.927 8H21v5a4 4 0 0 1-4 4h-.933a6.67 6.67 0 0 1-2.583 3H20v2H3v-2h3.506a6.597 6.597 0 0 1-3.084-5.039v-.007L2.926 8Zm13.648 7H17a2 2 0 0 0 2-2v-3h-2.068l-.353 4.954v.007a5.05 5.05 0 0 1-.004.039Zm-1.648-5H5.074l.341 4.797A4.603 4.603 0 0 0 10 19c2.059 0 3.836-1.38 4.412-3.289c.09-.295.145-.587.173-.914L14.928 10Z"/>`),
		g.Group(children),
	)
}