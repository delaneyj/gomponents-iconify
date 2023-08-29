package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func V(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m.59 15.956l171.09 473.93a29.099 29.099 0 0 0 27.37 19.218h123.777c6.73 0 12.737-4.222 15.016-10.555l173.57-482.481c2.43-6.754-3.026-13.73-10.166-12.998L363.502 17.186a24.03 24.03 0 0 0-20.251 16.021L256.817 282.08L169.795 32.097a24.03 24.03 0 0 0-20.43-16.023L10.679 2.94C3.56 2.267-1.838 9.231.59 15.956z"/>`),
		g.Group(children),
	)
}