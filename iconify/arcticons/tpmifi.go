package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tpmifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.96 12.618A20.053 20.053 0 0 1 24 4.5h0a20.053 20.053 0 0 1 16.04 8.102m-26.322 4.302A12.906 12.906 0 0 1 24 11.533a12.906 12.906 0 0 1 10.306 5.387m-15.241 4.03a6.185 6.185 0 0 1 9.87 0M12.958 43.476V21.521l10.99 13.348l10.99-13.315V43.5"/>`),
		g.Group(children),
	)
}