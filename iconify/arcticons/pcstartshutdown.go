package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pcstartshutdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="14" cy="14" r="8.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34" cy="14" r="8.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="14" cy="34" r="8.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34" cy="34" r="8.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.5 17.5l7-7m-7 0l7 7M12 9.968a4.5 4.5 0 1 0 4 0M14 8.5V13M9.5 34a4.5 4.5 0 1 0 1.318-3.182"/><circle cx="34" cy="38.557" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.234 31.266a2.766 2.766 0 0 1 5.532 0a2.477 2.477 0 0 1-.81 1.955c-.573.46-1.966 1.212-1.966 2.382v.377m-22.466-7.788l-.706 2.626l2.626.706"/>`),
		g.Group(children),
	)
}