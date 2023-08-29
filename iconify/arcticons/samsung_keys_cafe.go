package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungKeysCafe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.875 12.588A2.088 2.088 0 0 0 34 11.512a2.088 2.088 0 0 0-3.735 1.822C30.955 15.322 34 17.5 34 17.5s3.045-2.18 3.735-4.165c.09-.232.14-.483.14-.747h0Z"/><circle cx="24" cy="24" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="14" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="14" cy="24" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="14" cy="14" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34" cy="24" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="27" height="7" x="10.5" y="30.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.5" ry="3.5"/>`),
		g.Group(children),
	)
}