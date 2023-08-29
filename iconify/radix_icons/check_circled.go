package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckCircled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877ZM1.827 7.5a5.673 5.673 0 1 1 11.346 0a5.673 5.673 0 0 1-11.346 0Zm8.332-1.962a.5.5 0 0 0-.818-.576L6.52 8.972L5.357 7.787a.5.5 0 0 0-.714.7L6.227 10.1a.5.5 0 0 0 .765-.062l3.167-4.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}