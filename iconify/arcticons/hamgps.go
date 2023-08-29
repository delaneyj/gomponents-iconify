package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hamgps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="21.8" ry="12.53" transform="rotate(-30 24.01 23.994)"/><ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="15.28" ry="8.45" transform="rotate(-30 24.01 23.994)"/><ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="8.56" ry="4.37" transform="rotate(-30 24.01 23.994)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.74 13.15l12.52 21.7m-25.14.05l37.76-21.8"/><circle cx="33.36" cy="22.32" r=".75" fill="currentColor"/><circle cx="15.61" cy="36.36" r=".75" fill="currentColor"/><circle cx="25.12" cy="20.95" r=".75" fill="currentColor"/><circle cx="35.61" cy="11.05" r=".75" fill="currentColor"/><circle cx="39.48" cy="21.7" r=".75" fill="currentColor"/><circle cx="39.48" cy="24.47" r=".75" fill="currentColor"/><circle cx="7.23" cy="30.25" r=".75" fill="currentColor"/><circle cx="11.43" cy="20.95" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}