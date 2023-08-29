package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Graph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 145.532v908.936h1200V145.532H0zM147.51 291.87h904.98v616.26H147.51V291.87zm769.922 48.779L763.184 608.13L690.82 457.251l-45.41-94.775l-56.689 88.403L461.938 648.34l-52.075-54.712l-37.866-39.771l-42.114 35.229l-148.902 124.293l75.073 89.941l106.787-89.136l65.698 68.994l51.489 54.053l40.283-62.769l110.303-171.899l75 156.519l48.047 100.269l55.518-96.387l209.766-363.794l-101.513-58.521z"/>`),
		g.Group(children),
	)
}