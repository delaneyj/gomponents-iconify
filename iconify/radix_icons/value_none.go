package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ValueNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 0 0-5.023 10.94l-.83.83a.5.5 0 1 0 .707.707l.83-.83a6.623 6.623 0 0 0 9.34-9.34l.83-.83a.5.5 0 1 0-.708-.708l-.83.83A6.597 6.597 0 0 0 7.5.878Zm3.642 2.274a5.673 5.673 0 0 0-7.992 7.992l7.992-7.992Zm-7.284 8.698a5.673 5.673 0 0 0 7.992-7.992L3.857 11.85Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}