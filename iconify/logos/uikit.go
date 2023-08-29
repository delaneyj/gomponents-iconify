package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uikit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#2396F3" d="M174.656 28.438L125.717 0L74.646 31.723l49.621 27.414l50.389-30.699Zm23.403 12.245L147.52 71.424l57.28 33.109V192l-77.248 43.904L51.2 192v-68.267L0 98.304v123.563l125.717 74.666L256 222.422V74.155l-57.941-33.472Z"/>`),
		g.Group(children),
	)
}