package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.503 9.711a13.703 13.703 0 0 1 20.994 0M18.73 17.887a50.083 50.083 0 0 0-14.23.527v24.77c5.675-2.187 12.167-1.673 19.5-.527m5.27-24.77a50.083 50.083 0 0 1 14.23.527v24.77c-5.675-2.187-12.167-1.673-19.5-.527"/><circle cx="24" cy="18.413" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.752 14.115a6.851 6.851 0 0 1 10.496 0M24 23.684v18.973m3.162-11.595h12.122m-31.094 0h12.12m-6.06-6.324v12.648"/>`),
		g.Group(children),
	)
}