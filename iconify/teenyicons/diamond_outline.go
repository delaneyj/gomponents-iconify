package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5 14.5l-.395.307a.5.5 0 0 0 .79 0L7.5 14.5Zm-7-9l-.429-.257a.5.5 0 0 0 .034.564L.5 5.5Zm3-5V0h-.283L3.07.243L3.5.5Zm8 0l.429-.257L11.783 0H11.5v.5Zm3 5l.395.307a.5.5 0 0 0 .034-.564L14.5 5.5Zm-6.605 8.693l-7-9l-.79.614l7 9l.79-.614ZM.929 5.757l3-5L3.07.243l-3 5l.858.514ZM3.5 1h8V0h-8v1Zm7.571-.243l3 5l.858-.514l-3-5l-.858.514Zm3.034 4.436l-7 9l.79.614l7-9l-.79-.614Z"/>`),
		g.Group(children),
	)
}