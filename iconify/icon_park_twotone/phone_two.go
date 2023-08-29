package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPhoneTwo0"><g fill="#555" stroke="#fff"><path stroke-width="4" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke-linejoin="round" stroke-width="3.918" d="M19.44 14c.444 0 .854.241 1.07.63l1.496 2.695c.196.353.205.78.024 1.14L21 21s.252 2.252 2 4c1.748 1.748 4 2 4 2l2.527-1.038c.36-.18.788-.17 1.141.025l2.703 1.503c.388.216.629.625.629 1.07v3.103c0 1.58-1.468 2.721-2.965 2.216c-3.076-1.038-7.85-3.013-10.875-6.04c-3.026-3.025-5.001-7.798-6.039-10.874c-.505-1.497.636-2.965 2.216-2.965h3.103Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPhoneTwo0)"/>`),
		g.Group(children),
	)
}