package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rtrnettest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24.421" r="10.848" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.682 25.079l3.532-4.888l-4.852 3.568a.908.908 0 0 0-.25.892a1.07 1.07 0 0 0 .678.678a.953.953 0 0 0 .892-.25ZM9.704 12.648v26.068"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.718 28.625a8.41 8.41 0 1 1 14.565 0M9.685 38.716h25.247m-3.364-28.591H6.341m-.001 0V32.83m5.869 9.25h26.927M16.432 6.761H42.5m0 0v24.387M39.136 42.08V21.057"/><circle cx="35.773" cy="38.716" r=".841" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="39.136" cy="20.216" r=".841" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="6.341" cy="33.67" r=".841" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.591" cy="6.761" r=".841" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.477 27.784h5.046m-4.205 1.682h3.458"/>`),
		g.Group(children),
	)
}