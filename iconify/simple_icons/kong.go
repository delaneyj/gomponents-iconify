package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.88 18.96h4.405l2.286 2.876l-.393.979h-5.69l.139-.979l-1.341-2.117l.594-.759Zm3.152-12.632l2.36-.004L24 18.97l-.824 3.845h-4.547l.283-1.083L9 9.912l2.032-3.584Zm4.17-5.144l4.932 3.876l-.632.651l.855 1.191v1.273l-2.458 2.004l-4.135-4.884h-2.407l.969-1.777l2.876-2.334ZM4.852 13.597l3.44-2.989l4.565 5.494l-1.296 2.012h-4.21l-2.912 3.822l-.665.879H0v-4.689l3.517-4.529h1.335Z"/>`),
		g.Group(children),
	)
}