package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Documents(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.364 4.51a1.994 1.994 0 0 0-1.945 1.994v35.002a1.994 1.994 0 0 0 1.944 1.994h27.224a1.994 1.994 0 0 0 1.994-1.994V14.472h-7.977a1.994 1.994 0 0 1-1.945-1.995V4.5Zm19.205 0l9.962 9.962m-23.693 8.456h16.274M15.838 34.994h16.274m-16.274-6.033h16.274"/>`),
		g.Group(children),
	)
}