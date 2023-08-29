package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.167 42.941V25.926M22.611 44.5V27.103m-2.778 13.791V25.926m5.556 14.968V27.103"/><circle cx="17.876" cy="24.569" r="2.317" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.124" cy="24.569" r="2.317" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.124" cy="24.569" r=".75" fill="currentColor"/><circle cx="17.876" cy="24.569" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.25 22.256a24.873 24.873 0 0 0 2.01-9.797v-.699A8.26 8.26 0 0 0 24 3.5h0m-4.213 22.38a4.056 4.056 0 0 0 2.903 1.223h2.62a4.056 4.056 0 0 0 2.903-1.223M24 3.5h0a8.26 8.26 0 0 0-8.26 8.26v.699a24.874 24.874 0 0 0 2.01 9.797"/>`),
		g.Group(children),
	)
}