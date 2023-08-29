package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.415 2.395a1.99 1.99 0 0 1 1.17 0l2.986.918a16.722 16.722 0 0 1 4.39 2.089c1.054.705.555 2.348-.713 2.348H4.752c-1.268 0-1.767-1.643-.714-2.348a16.721 16.721 0 0 1 4.391-2.09l2.986-.917Zm.73 1.434a.49.49 0 0 0-.29 0l-2.985.918A15.22 15.22 0 0 0 5.5 6.25h13a15.22 15.22 0 0 0-3.37-1.503l-2.986-.918Z" clip-rule="evenodd"/><path fill="currentColor" d="M4.25 21a.75.75 0 0 1 .75-.75h14a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1-.75-.75Zm2-4a.75.75 0 0 0 1.5 0v-6a.75.75 0 0 0-1.5 0v6Zm5.75.75a.75.75 0 0 1-.75-.75v-6a.75.75 0 0 1 1.5 0v6a.75.75 0 0 1-.75.75Zm4.25-.75a.75.75 0 0 0 1.5 0v-6a.75.75 0 0 0-1.5 0v6Z"/>`),
		g.Group(children),
	)
}