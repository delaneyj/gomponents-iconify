package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurgicalKnife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M25.274 23.077h.967v.922h-.967zm1.942 0h.967v.922h-.967zm-7.762 0h.967v.922h-.967zm3.878 0h.967v.922h-.967zm-1.934 0h.967v.922h-.967zm-14.208 0c-.908.351-2.002.67-3.129.9l-.138.024h14.729v-.923zM18.344 10.14l7.898-7.898l-2.24-2.24L0 24a20.794 20.794 0 0 0 18.299-13.715zm5.92-9.239l.659.659L17.505 9l-.661-.659z"/>`),
		g.Group(children),
	)
}