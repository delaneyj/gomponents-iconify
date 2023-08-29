package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MightyDoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.846 32.221l-7.278 5.465V18.814h7.278v13.407zm1.918-1.277l4.913 3.689l2.364-1.952V18.814h-7.277v12.13zm16.472 0l-4.913 3.689l-2.364-1.952V18.814h7.277v12.13zM6.568 18.814H4.5m36.932 18.872V18.814h-1.273l-2.421 15.452l-2.312-15.452h-1.272v13.407m7.278-13.407H43.5m-27.509-8.478v5.978m-8.682-.006v-5.972l2.989 5.978l2.989-5.969v5.969m17.358-5.978h3.961m-1.98 5.978v-5.978m-8.044 0v5.978m3.961-5.978v5.978m-3.961-3h3.961m12.148-3l-1.98 2.99l-1.981-2.99m1.981 5.979v-2.989m-16.375-.988a1.98 1.98 0 0 0-1.98-1.98h0a1.98 1.98 0 0 0-1.98 1.98v2.018a1.98 1.98 0 0 0 1.98 1.98h0a1.98 1.98 0 0 0 1.98-1.98h-1.98"/>`),
		g.Group(children),
	)
}