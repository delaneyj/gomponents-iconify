package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Documents(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 6v5.5h-8v-10H9v2.75C9 5.216 9.784 6 10.75 6h2.75Zm-3-4.379L13.379 4.5H10.75a.25.25 0 0 1-.25-.25V1.621ZM5 0a1 1 0 0 0-1 1v2H2.25C1.56 3 1 3.56 1 4.25v10.5c0 .69.56 1.25 1.25 1.25h8.25c.69 0 1.25-.56 1.25-1.25v-1.67a.756.756 0 0 0-.004-.08H14a1 1 0 0 0 1-1V4.414a1 1 0 0 0-.293-.707L11.293.293A1 1 0 0 0 10.586 0H5ZM4 12V4.5H2.5v10h7.75v-1.42c0-.027.001-.054.004-.08H5a1 1 0 0 1-1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}