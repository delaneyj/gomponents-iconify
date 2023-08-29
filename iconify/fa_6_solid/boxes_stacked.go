package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxesStacked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M248 0h-40c-26.5 0-48 21.5-48 48v112c0 35.3 28.7 64 64 64h128c35.3 0 64-28.7 64-64V48c0-26.5-21.5-48-48-48h-40v80c0 8.8-7.2 16-16 16h-48c-8.8 0-16-7.2-16-16V0zM64 256c-35.3 0-64 28.7-64 64v128c0 35.3 28.7 64 64 64h160c35.3 0 64-28.7 64-64V320c0-35.3-28.7-64-64-64h-40v80c0 8.8-7.2 16-16 16h-48c-8.8 0-16-7.2-16-16v-80H64zm288 256h160c35.3 0 64-28.7 64-64V320c0-35.3-28.7-64-64-64h-40v80c0 8.8-7.2 16-16 16h-48c-8.8 0-16-7.2-16-16v-80h-40c-15 0-28.8 5.1-39.7 13.8c4.9 10.4 7.7 22 7.7 34.2v160c0 12.2-2.8 23.8-7.7 34.2C323.2 506.9 337 512 352 512z"/>`),
		g.Group(children),
	)
}