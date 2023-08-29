package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.511 2.634a1.74 1.74 0 0 0-1.022 0l-2.986.918A16.471 16.471 0 0 0 4.178 5.61c-.848.567-.446 1.89.574 1.89h14.496c1.02 0 1.422-1.323.575-1.89a16.472 16.472 0 0 0-4.326-2.058l-2.986-.918ZM4.25 21a.75.75 0 0 1 .75-.75h14a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1-.75-.75Zm2-4a.75.75 0 0 0 1.5 0v-6a.75.75 0 0 0-1.5 0v6Zm5.75.75a.75.75 0 0 1-.75-.75v-6a.75.75 0 0 1 1.5 0v6a.75.75 0 0 1-.75.75Zm4.25-.75a.75.75 0 0 0 1.5 0v-6a.75.75 0 0 0-1.5 0v6Z"/>`),
		g.Group(children),
	)
}