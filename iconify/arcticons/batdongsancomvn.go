package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Batdongsancomvn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.1 7.29a2.68 2.68 0 0 1 3.8 0l15.194 15.194a2.68 2.68 0 0 1 0 3.798l-5.698 5.698a2.68 2.68 0 0 1-3.799 0m-15.194 0l5.698-5.698a2.68 2.68 0 0 1 3.798 0l5.698 5.698m-15.195 0a2.68 2.68 0 0 1-3.798 0l-5.698-5.698a2.68 2.68 0 0 1 0-3.798L22.101 7.289a2.68 2.68 0 0 1 3.798 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.1 16.02a2.68 2.68 0 0 1 3.8 0l15.194 15.194a2.68 2.68 0 0 1 0 3.799l-5.698 5.698a2.68 2.68 0 0 1-3.799 0m-15.194 0l5.698-5.698a2.68 2.68 0 0 1 3.798 0l5.698 5.698m-15.195 0a2.68 2.68 0 0 1-3.798 0l-5.698-5.698a2.68 2.68 0 0 1 0-3.799L22.101 16.02a2.68 2.68 0 0 1 3.798 0m3.678-5.052V9.42c0-2.686 3.799-2.686 3.799 0v5.346"/>`),
		g.Group(children),
	)
}