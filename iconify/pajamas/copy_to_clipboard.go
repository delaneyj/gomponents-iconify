package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyToClipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM9.45 2a2.5 2.5 0 0 0-4.9 0H3a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h2v-1.5H3.5v-9h1V5h5V3.5h1V7H12V3a1 1 0 0 0-1-1H9.45ZM7.5 9.5h1.25a.75.75 0 0 0 0-1.5h-1.5C6.56 8 6 8.56 6 9.25v1.5a.75.75 0 0 0 1.5 0V9.5Zm1.25 5H7.5v-1.25a.75.75 0 0 0-1.5 0v1.5c0 .69.56 1.25 1.25 1.25h1.5a.75.75 0 0 0 0-1.5Zm3.75-5h-1.25a.75.75 0 0 1 0-1.5h1.5c.69 0 1.25.56 1.25 1.25v1.5a.75.75 0 0 1-1.5 0V9.5Zm-1.25 5h1.25v-1.25a.75.75 0 0 1 1.5 0v1.5c0 .69-.56 1.25-1.25 1.25h-1.5a.75.75 0 0 1 0-1.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}