package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 7v18h2V7zm4 0v18h6V7zm8 0v18h2V7zm4 0v18h4V7zm6 0v18h2V7zm4 0v18h2V7z"/>`),
		g.Group(children),
	)
}