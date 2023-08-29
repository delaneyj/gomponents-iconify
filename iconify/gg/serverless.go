package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Serverless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.787 6H5v3h5.695l1.092-3Zm-1.82 5H5v3h3.875l1.092-3Zm1.037 3l1.092-3H20v3h-8.996Zm-2.856 2H5v3h2.056l1.092-3Zm1.036 3l1.092-3H20v3H9.184Zm3.64-10l1.092-3H20v3h-7.176Z"/>`),
		g.Group(children),
	)
}