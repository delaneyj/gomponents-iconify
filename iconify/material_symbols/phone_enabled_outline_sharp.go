package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneEnabledOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.05 21H3v-5.875L8 14.1l2.9 2.9q1-.575 1.863-1.238t1.587-1.387q.775-.75 1.45-1.625t1.225-1.85l-2.85-2.875L15.1 3H21v1.05q0 3.15-1.35 6.2T15.8 15.8q-2.5 2.5-5.563 3.85T4.05 21ZM17.975 9q.375-.9.625-1.888T18.975 5H16.75l-.425 2.35L17.975 9Zm-8.95 8.95L7.35 16.275L5 16.75v2.2q1-.075 2.013-.325t2.012-.675ZM17.975 9Zm-8.95 8.95Z"/>`),
		g.Group(children),
	)
}