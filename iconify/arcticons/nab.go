package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.99 24l10.09-4.773l-2.559-10.864l10.022 4.913l6.898-8.775L31.849 15.4l11.161-.078L35.99 24l7.02 8.678l-11.161-.078l-2.408 10.899l-6.898-8.775l-10.022 4.913l2.559-10.864L4.99 24zm26.859-8.6L22.27 26.344m15.441-10.982l-9.777 11.143"/>`),
		g.Group(children),
	)
}