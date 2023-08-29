package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Breathly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="33.75" cy="7.131" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.13" cy="11.67" r=".75" fill="currentColor"/><circle cx="40.9" cy="14.269" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.361" cy="16.889" r=".75" fill="currentColor"/><circle cx="43.5" cy="24.019" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="38.259" cy="24.019" r=".75" fill="currentColor"/><circle cx="40.888" cy="33.769" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.349" cy="31.148" r=".75" fill="currentColor"/><circle cx="33.75" cy="40.906" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.13" cy="36.367" r=".75" fill="currentColor"/><circle cx="24" cy="43.5" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="38.259" r=".75" fill="currentColor"/><circle cx="14.25" cy="40.906" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="16.87" cy="36.367" r=".75" fill="currentColor"/><circle cx="7.112" cy="33.769" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.651" cy="31.148" r=".75" fill="currentColor"/><circle cx="4.5" cy="24.019" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="9.741" cy="24.019" r=".75" fill="currentColor"/><circle cx="7.112" cy="14.269" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.651" cy="16.889" r=".75" fill="currentColor"/><circle cx="14.25" cy="7.131" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="16.87" cy="11.67" r=".75" fill="currentColor"/><circle cx="24" cy="4.5" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="9.741" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}