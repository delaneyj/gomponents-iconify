package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spdx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0v24h8.222l2.089-2.373l2.09-2.374V13.2h6.577l2.51-2.488L24 8.223V0H12zm5.2 5.2h13.791L12.2 12c-3.735 3.74-6.838 6.8-6.896 6.8c-.057 0-.104-3.06-.104-6.8zm8.4 8.8v10H24V14h-5.2z"/>`),
		g.Group(children),
	)
}