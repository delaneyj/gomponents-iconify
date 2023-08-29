package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 13.5a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8.91l5.88 5.88a3 3 0 0 0 4.24 0l3.59-3.58a1 1 0 0 0-1.42-1.42l-3.58 3.59a1 1 0 0 1-1.42 0L5.41 7.5H17a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1Zm0-11a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1Zm-.2 7a.64.64 0 0 0-.18.06a.76.76 0 0 0-.18.09l-.15.12a1.05 1.05 0 0 0-.29.71a1.23 1.23 0 0 0 0 .19a.6.6 0 0 0 .06.19a.76.76 0 0 0 .09.18a1.58 1.58 0 0 0 .12.15a1 1 0 0 0 1.42 0l.12-.15a.76.76 0 0 0 .09-.18a.6.6 0 0 0 .06-.19a1.23 1.23 0 0 0 0-.19a1 1 0 0 0-1.2-1Z"/>`),
		g.Group(children),
	)
}