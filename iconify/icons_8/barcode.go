package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 7v18h2V7H2zm4 0v18h6V7H6zm8 0v18h2V7h-2zm4 0v18h4V7h-4zm6 0v18h2V7h-2zm4 0v18h2V7h-2z"/>`),
		g.Group(children),
	)
}