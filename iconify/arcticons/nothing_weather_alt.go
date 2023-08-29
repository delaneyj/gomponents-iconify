package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NothingWeatherAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="34.5" cy="35.5" r=".75" fill="currentColor"/><circle cx="31.5" cy="38.5" r=".75" fill="currentColor"/><circle cx="25.5" cy="35.5" r=".75" fill="currentColor"/><circle cx="22.5" cy="38.5" r=".75" fill="currentColor"/><circle cx="16.5" cy="35.5" r=".75" fill="currentColor"/><circle cx="13.5" cy="38.5" r=".75" fill="currentColor"/><circle cx="31.5" cy="29.5" r=".75" fill="currentColor"/><circle cx="34.5" cy="29.5" r=".75" fill="currentColor"/><circle cx="16.5" cy="29.5" r=".75" fill="currentColor"/><circle cx="19.5" cy="29.5" r=".75" fill="currentColor"/><circle cx="22.5" cy="29.5" r=".75" fill="currentColor"/><circle cx="25.5" cy="29.5" r=".75" fill="currentColor"/><circle cx="28.5" cy="29.5" r=".75" fill="currentColor"/><circle cx="34.5" cy="14.5" r=".75" fill="currentColor"/><circle cx="37.5" cy="17.5" r=".75" fill="currentColor"/><circle cx="28.5" cy="14.5" r=".75" fill="currentColor"/><circle cx="31.5" cy="14.5" r=".75" fill="currentColor"/><circle cx="16.5" cy="17.5" r=".75" fill="currentColor"/><circle cx="19.5" cy="17.5" r=".75" fill="currentColor"/><circle cx="22.25" cy="20.25" r=".75" fill="currentColor"/><circle cx="25.5" cy="17.5" r=".75" fill="currentColor"/><circle cx="13.5" cy="20.5" r=".75" fill="currentColor"/><circle cx="37.5" cy="29.5" r=".75" fill="currentColor"/><circle cx="40.5" cy="23.5" r=".75" fill="currentColor"/><circle cx="40.5" cy="26.5" r=".75" fill="currentColor"/><circle cx="40.5" cy="20.5" r=".75" fill="currentColor"/><circle cx="10.5" cy="23.5" r=".75" fill="currentColor"/><circle cx="10.5" cy="26.5" r=".75" fill="currentColor"/><circle cx="13.5" cy="29.5" r=".75" fill="currentColor"/><circle cx="10.25" cy="12.25" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}