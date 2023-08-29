package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnavailableLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20.25A8.25 8.25 0 0 1 3.75 12h-1.5A9.75 9.75 0 0 0 12 21.75v-1.5Zm0-16.5A8.25 8.25 0 0 1 20.25 12h1.5A9.75 9.75 0 0 0 12 2.25v1.5ZM3.75 12c0-2.278.923-4.34 2.416-5.834l-1.06-1.06A9.722 9.722 0 0 0 2.25 12h1.5Zm2.416-5.834A8.222 8.222 0 0 1 12 3.75v-1.5a9.722 9.722 0 0 0-6.894 2.856l1.06 1.06Zm-1.06 0l12.728 12.728l1.06-1.06L6.166 5.106l-1.06 1.06ZM20.25 12c0 2.278-.923 4.34-2.416 5.834l1.06 1.06A9.722 9.722 0 0 0 21.75 12h-1.5Zm-2.416 5.834A8.222 8.222 0 0 1 12 20.25v1.5a9.722 9.722 0 0 0 6.894-2.856l-1.06-1.06Z"/>`),
		g.Group(children),
	)
}