package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m10.72 2.598l-.27.902L10 5h3.989a2 2 0 0 1 1.932 2.517l-1.19 4.448a4 4 0 0 1-4.88 2.835L3 13H0V5h3L7.39.61c1.47-1.47 3.928-.002 3.33 1.988ZM3 6.5H1.5v5H3v-5Zm7.232 6.85L4.5 11.843V5.621L8.451 1.67a.466.466 0 0 1 .296-.155a.533.533 0 0 1 .314.08c.11.065.183.154.22.238c.03.07.051.173.003.334l-.72 2.402l-.58 1.931h6.005a.5.5 0 0 1 .483.63l-1.19 4.447a2.5 2.5 0 0 1-3.05 1.773Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}