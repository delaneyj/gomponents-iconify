package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Powerschool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.81 23.752a5.458 5.458 0 1 0 0-10.916V8.668a9.626 9.626 0 0 1 0 19.252v4.168a13.794 13.794 0 0 0 0-27.588h-5.395v39h-5.01v-39H9.397v39"/>`),
		g.Group(children),
	)
}