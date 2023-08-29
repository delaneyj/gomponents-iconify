package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opensstpclient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="27.405" x="4.5" y="14.668" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.108"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.233 14.667a11.588 11.588 0 0 0-22.466 0"/><circle cx="37.176" cy="35.75" r="3.162" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37.176" cy="29.425" r="3.162" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.851" cy="35.75" r="3.162" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.851" cy="29.425" r="3.162" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}