package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoStencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M188.8 334.07h197.33L279.47 448H83.2ZM512 199H106.61L0 313h405.39ZM232.2 64h196.6L322.62 177.93H125.87Z"/>`),
		g.Group(children),
	)
}