package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MavVonatinfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" d="m5.787 26.871l.777 2.955M3.5 18.176l.778 2.955m37.935 5.74l-.777 2.954m3.064-11.65l-.778 2.955m-8.27 5.741h6.761m2.287-8.696H34.325m9.397 2.955h-8.274m-1.13 8.695h7.118m-17.938-2.954H5.788m-1.51-5.74h19.22M3.502 18.175h19.995m0 11.65H6.564m16.934-8.696a2.87 2.87 0 0 1 0 5.74m0-8.695a5.83 5.83 0 0 1 .842 11.588a5.701 5.701 0 0 1-.842.062m-1.288-14.98a9.247 9.247 0 1 1-5.817 14.982m-.01-11.656a9.137 9.137 0 0 1 5.727-3.321m-.082-2.93a12.2 12.2 0 0 1 3.464 24.153a12.03 12.03 0 0 1-12.252-6.031m.003-12.086a12.071 12.071 0 0 1 8.656-6.016"/>`),
		g.Group(children),
	)
}