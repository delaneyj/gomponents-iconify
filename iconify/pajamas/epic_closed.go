package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EpicClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m4 2.5l-2.5 6h13l-2.5-6H4Zm-4 8.25V8.569c-.01-.212.026-.431.115-.646l2.629-6.308A1 1 0 0 1 3.667 1h8.666a1 1 0 0 1 .923.615l2.629 6.308A1.5 1.5 0 0 1 14.5 10h-13v.75c0 .138.112.25.25.25h12a.75.75 0 0 1 0 1.5H3.5v.75c0 .138.112.25.25.25h8.5a.75.75 0 0 1 0 1.5h-8.5A1.75 1.75 0 0 1 2 13.25v-.75h-.25A1.75 1.75 0 0 1 0 10.75Zm10.28-5.72a.75.75 0 1 0-1.06-1.06L7.5 5.69l-.72-.72a.75.75 0 1 0-1.06 1.06l1.25 1.25a.75.75 0 0 0 1.06 0l2.25-2.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}