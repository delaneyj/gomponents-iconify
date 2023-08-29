package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Desktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 16.75H5A1.76 1.76 0 0 1 3.25 15V5A1.76 1.76 0 0 1 5 3.25h14A1.76 1.76 0 0 1 20.75 5v10A1.76 1.76 0 0 1 19 16.75Zm-14-12a.25.25 0 0 0-.25.25v10a.25.25 0 0 0 .25.25h14a.25.25 0 0 0 .25-.25V5a.25.25 0 0 0-.25-.25Z"/><path fill="currentColor" d="M15 20.75h-3a.76.76 0 0 1-.75-.75v-4a.75.75 0 0 1 1.5 0v3.25H15a.75.75 0 0 1 0 1.5Z"/><path fill="currentColor" d="M12 20.75H9a.75.75 0 0 1 0-1.5h3a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}