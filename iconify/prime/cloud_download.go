package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 18.25a.75.75 0 0 1 0-1.5c1.66 0 2.25-.83 2.25-3.18a3.57 3.57 0 0 0-3.25-3.25a3.34 3.34 0 0 0-1 .18a.74.74 0 0 1-1-.49a5.25 5.25 0 0 0-10.25 1.56c0 3.22.15 5.18 2.25 5.18a.75.75 0 0 1 0 1.5c-3.75 0-3.75-3.86-3.75-6.68a6.75 6.75 0 0 1 13-2.68a4.4 4.4 0 0 1 .8-.07a5.07 5.07 0 0 1 4.75 4.75c-.05 1.28-.05 4.68-3.8 4.68Z"/><path fill="currentColor" d="M12 19.18a.74.74 0 0 1-.53-.22l-2.83-2.83a.75.75 0 0 1 1.06-1.06l2.3 2.3l2.3-2.3a.75.75 0 0 1 1.06 1.06L12.53 19a.74.74 0 0 1-.53.18Z"/><path fill="currentColor" d="M12 19.18a.75.75 0 0 1-.75-.75v-6.36a.75.75 0 0 1 1.5 0v6.36a.75.75 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}