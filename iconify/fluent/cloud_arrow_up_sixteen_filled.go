package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudArrowUpSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.03 5.507a4 4 0 0 1 7.887-.323A5.5 5.5 0 0 0 5.207 12H4.25a3.25 3.25 0 0 1-.22-6.493ZM15 10.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-4.854-2.353l-2 2a.5.5 0 0 0 .708.707L10 9.707V12.5a.5.5 0 0 0 1 0V9.707l1.146 1.147a.5.5 0 0 0 .708-.708l-2-2A.499.499 0 0 0 10.503 8h-.006a.498.498 0 0 0-.348.144l-.003.003Z"/>`),
		g.Group(children),
	)
}