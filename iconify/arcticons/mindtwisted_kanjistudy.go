package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MindtwistedKanjistudy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.542 4.5c1.681 1.46 2.694 3.736 3.327 6.478M10.844 14.46c7.937-.468 19.022-6.182 27.994-6.02a8.64 8.64 0 0 1-3.326 5.952"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.343 9.971c1.199 3.344.558 6.253-1.663 9.499m7.441-.744c5.713.107 10.264-3.527 14.576-2.802a88.955 88.955 0 0 1-7.529 7.223C27 28.758 31.258 34.254 22.03 43.5a17.089 17.089 0 0 1-3.72-5.077"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.65 29.625c9.671-.73 20.282-5.252 28.67-5.253"/>`),
		g.Group(children),
	)
}