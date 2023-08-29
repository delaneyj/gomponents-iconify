package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeployedCodeUpdateSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17.125V6.875L12 1.7l9 5.175v5.375q-.85-.6-1.863-.925T17 11q-.425 0-.813.038t-.787.137L19 9.1V8.05l-1.075-.625L12 10.85L6.075 7.425L5 8.05V9.1l6 3.475V14.4q-.475.8-.738 1.7T10 18q0 1.05.288 2t.812 1.775q-.025-.025-.05-.025t-.05-.025l-8-4.6ZM17 21l3-3l-.7-.7l-1.8 1.8V15h-1v4.1l-1.8-1.8l-.7.7l3 3Zm0 2q-2.075 0-3.538-1.463T12 18q0-2.075 1.463-3.538T17 13q2.075 0 3.538 1.463T22 18q0 2.075-1.463 3.538T17 23Z"/>`),
		g.Group(children),
	)
}