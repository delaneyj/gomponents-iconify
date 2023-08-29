package medical_icon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IStairs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M1.648 56.572h11.205V45.168h11.134V33.686h11.308v-11.27h11.233V11.222h16.677v5.711H52.118V28.17H40.872v11.397H29.65v11.336H18.383v11.432H1.648z"/>`),
		g.Group(children),
	)
}