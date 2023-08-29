package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightTracks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M117 195h193q31 0 53-22t22-53t-22-53.5T309.5 44T256 66.5T234 120v33h-42v-33q0-49 34.5-83.5t83-34.5t83 34.5t34.5 83t-34.5 83T310 237H117q-31 0-53 22t-22 53t22 53t53 22t53-22t22-53v-34h42v34q0 48-34 82.5T117.5 429t-83-34.5t-34.5-83T34.5 229t82.5-34z"/>`),
		g.Group(children),
	)
}