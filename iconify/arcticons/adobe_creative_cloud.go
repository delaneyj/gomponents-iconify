package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeCreativeCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4" ry="4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.28 33.615c-2.97.108-4.786.277-7.286-1.37a7.97 7.97 0 0 1 5.784-14.455c2.915.503 5.133 3.077 7.601 6.144m-5.193-7.626a9.747 9.747 0 0 1 2.009-1.217a9.571 9.571 0 0 1 13.304 8.753a9.527 9.527 0 0 1-.575 3.382a8.818 8.818 0 0 1-1.53 2.77a9.463 9.463 0 0 1-9.05 3.482c-3.597-.603-6.445-3.928-9.67-7.448"/>`),
		g.Group(children),
	)
}