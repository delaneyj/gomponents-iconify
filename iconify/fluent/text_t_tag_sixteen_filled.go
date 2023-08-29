package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextTTagSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.001 4.75v-1a.75.75 0 0 0-.75-.75h-6.5a.75.75 0 0 0-.75.75v1a.75.75 0 0 0 1.5 0V4.5h1.75v7h-.25a.75.75 0 0 0 0 1.5h2a.75.75 0 0 0 0-1.5h-.25v-7h1.75v.25a.75.75 0 0 0 1.5 0ZM4.249 6.689a.75.75 0 0 0-1.06.062l-2 2.25a.748.748 0 0 0 0 .996l2 2.25a.747.747 0 0 0 1.06.063a.749.749 0 0 0 .062-1.059L2.754 9.499L4.31 7.747a.75.75 0 0 0-.062-1.058Zm8.563.063l2 2.25v.001a.748.748 0 0 1 0 .996l-2 2.25a.75.75 0 1 1-1.121-.996l1.557-1.752l-1.557-1.752a.749.749 0 1 1 1.121-.997Z"/>`),
		g.Group(children),
	)
}