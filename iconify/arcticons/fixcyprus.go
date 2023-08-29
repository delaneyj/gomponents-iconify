package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fixcyprus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.269 43.5h31.462m-3.479 0l-9.191-39h-5.703l-9.19 39m22.034-6.807H14.217m18.145-8.462H16.058m14.464-8.462H17.898m10.784-8.462h-8.944"/>`),
		g.Group(children),
	)
}