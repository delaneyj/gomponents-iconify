package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Imagingedgemobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.05" d="M27.25 4.5h6.5V11h6.5v32.5H7.75V11h19.5Zm0 6.5v6.51h-13V37h19.5V11Zm0 0"/>`),
		g.Group(children),
	)
}