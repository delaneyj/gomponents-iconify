package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nethack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.587 14.55a12.35 12.35 0 0 0 5.913-.885L39.41 8.36l-1.823 1.05m-2.021 21.26a13.58 13.58 0 0 0 5.502 5.101l-3.095 1.659s.443 2.708 1.603 3.647c-3.537.609-7.375-.755-8.44-4.799"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.495 32.167s1.765 1.283.66 2.223a1.886 1.886 0 0 1-2.34-.06"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 42.072c-7.342-3.703-13.587-14.535-13.587-14.535V5.928l3.703 3.703l3.675-3.675H30.21l3.675 3.675l3.703-3.703v21.609S31.342 38.369 24 42.072Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.413 14.55a12.35 12.35 0 0 1-5.913-.885L8.59 8.36l1.824 1.05m2.02 21.26a13.58 13.58 0 0 1-5.502 5.101l3.095 1.659s-.443 2.708-1.603 3.647c3.537.609 7.375-.755 8.44-4.799"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.505 32.167s-1.765 1.283-.66 2.223a1.886 1.886 0 0 0 2.34-.06M24 5.956v36.116"/>`),
		g.Group(children),
	)
}