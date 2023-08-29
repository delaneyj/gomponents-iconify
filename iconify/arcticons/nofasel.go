package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nofasel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 38.414V9.586l19.099 28.828V9.586M28.086 24h9.369m-9.369 14.414V9.586H42.5"/>`),
		g.Group(children),
	)
}