package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitnex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="12.27" cy="39.86" r="3.64" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.27" cy="12.4" r="3.64" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.73" cy="12.4" r="3.64" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.73" cy="39.86" r="3.64" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="26.11" r="3.64" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.27 8.76V4.5m0 11.54v20.18m23.46 0V16.04M15.91 12.4H19a5 5 0 0 1 5 5v5.07m0 7.28v5.11a5 5 0 0 0 5 5h3.09"/>`),
		g.Group(children),
	)
}