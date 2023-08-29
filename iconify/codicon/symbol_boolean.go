package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolBoolean(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m1 3.5l.5-.5h13l.5.5v9l-.5.5h-13l-.5-.5v-9zM14 4H8v3.493h-.5l-3.574-.005l2.09-2.09l-.707-.707l-2.955 2.955v.708l2.955 2.955l.707-.707l-2.114-2.114l3.996.005H8v-.986l3.907.005l-2.114-2.114l.707-.707l2.956 2.955v.708L10.5 11.309l-.707-.707l2.09-2.09L8 8.507V12h6V4z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}