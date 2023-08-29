package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeployedCodeUpdateOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 11.425Zm-9 5.7V6.875L12 1.7l9 5.175v5.4q-.45-.325-.95-.562T19 11.3V9.1l-3.6 2.075q-1.4.35-2.538 1.188T11 14.4v-1.825L5 9.1v6.85l5.05 2.925q.1.8.375 1.538t.7 1.387q-.05-.025-.063-.038T11 21.726l-8-4.6ZM12 4L6.075 7.425L12 10.85l5.925-3.425L12 4Zm5 17l3-3l-.7-.7l-1.8 1.8V15h-1v4.1l-1.8-1.8l-.7.7l3 3Zm0 2q-2.075 0-3.538-1.463T12 18q0-2.075 1.463-3.538T17 13q2.075 0 3.538 1.463T22 18q0 2.075-1.463 3.538T17 23Z"/>`),
		g.Group(children),
	)
}