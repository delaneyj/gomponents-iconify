package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardiogramENegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsCardiogramENegative0)"><path d="M14 9a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1V9Zm-3 21a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H12a1 1 0 0 1-1-1Zm1 4a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2H12Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM15 6a3 3 0 0 0-3 3H9a3 3 0 0 0-3 3v9h7.035l3.298-5.534l2 7.339L20.51 20h5.258a1 1 0 1 1 0 2H21.49l-4.034 5.195l-1.815-6.661L14.171 23H6v16a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V12a3 3 0 0 0-3-3h-3a3 3 0 0 0-3-3h-8Zm21 12a3 3 0 1 1 6 0v3h-6v-3Zm0 20V23h6v15l-3 4l-3-4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsCardiogramENegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}