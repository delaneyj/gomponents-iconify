package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SexFemale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M640 1664v-256h384v-256H640v-144q168-43 276-181t108-315q0-212-150-362T512 0T150 150T0 512q0 177 108 315t276 181v144H0v256h384v256h256zM512 256q106 0 181 75t75 181t-75 181t-181 75t-181-75t-75-181t75-181t181-75z"/>`),
		g.Group(children),
	)
}