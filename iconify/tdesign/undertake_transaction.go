package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UndertakeTransaction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 3a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM10 5.5a4.5 4.5 0 0 1 6.5-4.032a4.5 4.5 0 1 1 0 8.064A4.5 4.5 0 0 1 10 5.5Zm8.25 2.488a2.5 2.5 0 1 0 0-4.975A4.48 4.48 0 0 1 19 5.5a4.48 4.48 0 0 1-.75 2.488ZM8.435 13.25a1.25 1.25 0 0 0-.885.364l-2.05 2.05V19.5h5.627l5.803-1.45l3.532-1.508a.555.555 0 0 0-.416-1.022l-.02.005L13.614 17H10v-2h3.125a.875.875 0 1 0 0-1.75h-4.69Zm7.552 1.152l3.552-.817a2.56 2.56 0 0 1 3.211 2.47a2.557 2.557 0 0 1-1.414 2.287l-.027.014l-3.74 1.595l-6.196 1.549H0v-7.25h4.086l2.052-2.052a3.25 3.25 0 0 1 2.3-.948h.002h-.002h4.687a2.875 2.875 0 0 1 2.862 3.152ZM3.5 16.25H2v3.25h1.5v-3.25Z"/>`),
		g.Group(children),
	)
}