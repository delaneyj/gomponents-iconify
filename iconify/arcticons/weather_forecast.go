package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherForecast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.692 31.156h-8.627c-9.018-.804-8.97-13.994.054-14.736c1.655-6.32 6.642-9.082 11.435-8.824a10.9 10.9 0 0 1 9.18 5.868c11.253 1.935 9.898 16.359 0 17.692h-9.187"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.048 8.713a6.685 6.685 0 1 0-9.44 8.698m15.542 7.953l-4.729 8.685l3.13 1.586l-1.35 6.559l5.22-7.73l-3.048-1.63Z"/><circle cx="17.457" cy="33.649" r=".75" fill="currentColor"/><circle cx="28.431" cy="39.625" r=".75" fill="currentColor"/><circle cx="30.495" cy="34.735" r=".75" fill="currentColor"/><circle cx="11.961" cy="34.842" r=".75" fill="currentColor"/><circle cx="17.898" cy="37.712" r=".75" fill="currentColor"/><circle cx="37.653" cy="34.036" r=".75" fill="currentColor"/><circle cx="34.493" cy="41.444" r=".75" fill="currentColor"/><circle cx="34.063" cy="36.917" r=".75" fill="currentColor"/><circle cx="14.126" cy="42.194" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}