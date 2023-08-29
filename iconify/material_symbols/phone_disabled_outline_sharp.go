package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneDisabledOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.975 9q.375-.9.625-1.888T18.975 5H16.75l-.425 2.35L17.975 9ZM5 18.95q1-.075 2.013-.325t2.012-.675L7.35 16.275L5 16.75v2.2Zm12.1-4.6l-1.45-1.45q.225-.275.738-1t.637-1l-2.85-2.875L15.1 3H21v1.05q0 2.725-.975 5.3t-2.925 5ZM4.05 21H3v-5.875L8 14.1l2.9 2.9q.725-.45 1.137-.738t.763-.562L1.4 4.3l1.4-1.4l18.4 18.4l-1.4 1.4l-5.55-5.55q-2.3 1.875-4.925 2.863T4.05 21ZM17.975 9Zm-8.95 8.95Z"/>`),
		g.Group(children),
	)
}