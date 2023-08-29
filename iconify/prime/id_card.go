package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 18.75H5A1.76 1.76 0 0 1 3.25 17V7A1.76 1.76 0 0 1 5 5.25h14A1.76 1.76 0 0 1 20.75 7v10A1.76 1.76 0 0 1 19 18.75Zm-14-12a.25.25 0 0 0-.25.25v10a.25.25 0 0 0 .25.25h14a.25.25 0 0 0 .25-.25V7a.25.25 0 0 0-.25-.25Z"/><path fill="currentColor" d="M9 11.75a2 2 0 1 1 2-2a2 2 0 0 1-2 2Zm0-2.5a.5.5 0 1 0 .5.5a.5.5 0 0 0-.5-.5Zm3 6.5a.76.76 0 0 1-.75-.75c0-.68-.17-1.25-2.25-1.25s-2.25.57-2.25 1.25a.75.75 0 0 1-1.5 0c0-2.75 2.82-2.75 3.75-2.75s3.75 0 3.75 2.75a.76.76 0 0 1-.75.75Zm5-5h-3a.75.75 0 0 1 0-1.5h3a.75.75 0 0 1 0 1.5Zm-1 3h-2a.75.75 0 0 1 0-1.5h2a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}