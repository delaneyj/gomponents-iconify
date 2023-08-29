package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Overseerr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 45.5c11.825 0 21.5-9.675 21.5-21.5S35.825 2.5 24 2.5S2.5 12.175 2.5 24S12.175 45.5 24 45.5Zm14.297-19.672c0 6.88-5.59 12.577-12.577 12.577s-12.47-5.697-12.47-12.578c0-1.29.215-2.472.537-3.655c.968 2.15 3.225 3.655 5.698 3.655a6.24 6.24 0 0 0 6.235-6.235c0-2.58-1.505-4.73-3.655-5.697c1.182-.322 2.365-.537 3.655-.537c6.987-.108 12.578 5.482 12.578 12.47h0Z"/>`),
		g.Group(children),
	)
}