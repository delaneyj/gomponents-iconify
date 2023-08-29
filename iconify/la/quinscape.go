package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quinscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4C9.373 4 4 9.373 4 16s5.373 12 12 12c1.315 0 2.568-.236 3.75-.629a6.51 6.51 0 0 1-.447-.38C15.253 26.883 12 23.574 12 19.5a7.5 7.5 0 0 1 7.5-7.5c4.075 0 7.384 3.253 7.49 7.303c.133.144.26.293.381.447c.393-1.182.629-2.435.629-3.75c0-6.627-5.373-12-12-12zm7 14a5 5 0 0 0 0 10a5 5 0 0 0 0-10z"/>`),
		g.Group(children),
	)
}