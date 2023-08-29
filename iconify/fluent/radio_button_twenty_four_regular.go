package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioButtonTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 22.002c5.524 0 10.002-4.478 10.002-10.001c0-5.524-4.478-10.002-10.002-10.002c-5.524 0-10.002 4.478-10.002 10.002c0 5.523 4.478 10.001 10.002 10.001Zm0-1.5A8.502 8.502 0 1 1 12 3.5a8.502 8.502 0 0 1 0 17.003Z"/>`),
		g.Group(children),
	)
}