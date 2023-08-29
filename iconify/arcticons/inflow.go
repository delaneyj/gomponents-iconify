package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.178 7.279c5.078-.543 11.078.506 15.349 3.341c5.758 3.823 6.74 11.398 1.764 16.285c-3.445 3.383-8.158 5.005-12.902 5.513c-1.127.12-5.52-.22-5.896 0c-2.755 1.621-.786 7.982-5.65 8.426c-2.65.241-5.166-2.444-6.649-4.288c-3.355-4.175-5.168-9.285-4.587-14.67C5.6 12.693 15.023 8.151 23.178 7.279m-6.384 23.164c-5.865-3.088-7.083-10.588-1.63-14.805c5.003-3.869 19.126-4.802 22.535 1.766c2.578 4.967-3.66 6.174-8.407 7.206c-1.717.373-3.238.724-4.077 1.22c-2.523 1.498-3.334 4.365-3.531 6.243"/>`),
		g.Group(children),
	)
}