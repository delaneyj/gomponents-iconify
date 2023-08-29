package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LadderOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M17 17L35 17"/><path d="M15 26L33 26"/><path d="M12 35L30 35"/><path d="M28.5652 43L38.3054 7.52959C38.655 6.25653 37.697 5 36.3768 5H22.2145C21.3374 5 20.5626 5.57158 20.3036 6.40968L9 43"/><path d="M17 18L21 42"/><path d="M35 18L39 42"/></g>`),
		g.Group(children),
	)
}