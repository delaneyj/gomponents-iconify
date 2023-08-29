package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Doom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.846 27.971l-7.278 5.465V14.564h7.278v13.407zm1.918-1.277l4.913 3.689l2.364-1.952V14.564h-7.277v12.13zm16.472 0l-4.913 3.689l-2.364-1.952V14.564h7.277v12.13zM6.568 14.564H4.5m36.932 18.872V14.564h-1.273l-2.421 15.452l-2.312-15.452h-1.272v13.407m7.278-13.407H43.5"/>`),
		g.Group(children),
	)
}