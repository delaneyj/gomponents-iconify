package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 1a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5zm6.994 5c-3.202 0-6.122 1.348-8.29 3.569a1 1 0 1 1-1.431-1.397C4.787 5.597 8.21 4 11.994 4c3.783 0 7.207 1.597 9.721 4.172a1 1 0 1 1-1.43 1.397C18.115 7.349 15.196 6 11.995 6zm0 5c-2.023 0-3.883.78-5.335 2.095a1 1 0 1 1-1.342-1.483C7.11 9.99 9.44 9 11.994 9c2.554 0 4.885.99 6.677 2.612a1 1 0 1 1-1.342 1.483C15.877 11.78 14.017 11 11.994 11zm0 5a4.3 4.3 0 0 0-2.511.814a1 1 0 1 1-1.162-1.628A6.3 6.3 0 0 1 11.994 14c1.363 0 2.627.44 3.673 1.186a1 1 0 1 1-1.162 1.628a4.3 4.3 0 0 0-2.51-.814zm-1 3a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2h-.01a1 1 0 0 1-1-1z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}