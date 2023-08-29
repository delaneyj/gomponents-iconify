package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jamesdsp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="8.14" cy="16.11" r="3.64" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="39.86" cy="16.11" r="3.64" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.14 12.47V8.4m0 11.35V39.6m31.72-27.13V8.4m0 11.35V39.6"/><circle cx="24" cy="31.89" r="3.64" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 35.53v4.07m0-11.35V8.4"/>`),
		g.Group(children),
	)
}