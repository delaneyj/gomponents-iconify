package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinuxSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5 8V7h1v1H5Zm4 0V7h1v1H9Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 6.5a6.5 6.5 0 0 1 13 0v6.514c0 .179.07.35.197.476l.657.656A.5.5 0 0 1 14.5 15H.5a.5.5 0 0 1-.354-.854l.657-.656A.673.673 0 0 0 1 13.014V6.5Zm3 0a1.5 1.5 0 1 1 3 0v1a1.5 1.5 0 1 1-3 0v-1Zm4 0a1.5 1.5 0 1 1 3 0v1a1.5 1.5 0 0 1-3 0v-1Zm-3.407 4.012a6.5 6.5 0 0 1 5.814 0l.249.125l-.095.095a4.329 4.329 0 0 1-6.122 0l-.095-.095l.249-.125Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}