package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.564 10.004c-.847 1.155-3.659 1.45-3.659 1.45s-2.876-.284-3.723-1.426C1.122 10.028 1 15.933 1 15.933h15.808c0 .001.319-5.929-4.244-5.929zm-.746-5.984c0 1.669-1.303 4.857-2.908 4.857C7.303 8.877 6 5.689 6 4.02C6 2.353 7.303 1 8.91 1c1.605.001 2.908 1.353 2.908 3.02z"/>`),
		g.Group(children),
	)
}