package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 4V2.5h4.697l1 1.5H1.5ZM0 4V2a1 1 0 0 1 1-1h5.465a1 1 0 0 1 .832.445l1.667 2.5l.034.055H15a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1V4Zm1.5 1.5v7h13v-7h-13Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}