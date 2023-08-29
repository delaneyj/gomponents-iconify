package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowOutlineDownLeftThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22.036 2.879a3 3 0 0 0-4.243 0L10.89 9.782L9.14 8.033c-2.12-2.121-5.754-.763-5.965 2.229l-1.167 16.53a3 3 0 0 0 3.204 3.204l16.533-1.168c2.993-.212 4.35-3.845 2.228-5.967l-1.75-1.748l6.903-6.902a3 3 0 0 0 0-4.242l-7.09-7.09Z"/>`),
		g.Group(children),
	)
}