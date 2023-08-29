package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openinwhatsapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="12.947" cy="7.878" r="3.378" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23.693" cy="7.878" r="3.378" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.442" cy="7.878" r="3.378" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="13.558" cy="18.627" r="3.378" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.307" cy="18.627" r="3.378" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.053" cy="18.627" r="3.378" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="13.558" cy="29.373" r="3.378" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.307" cy="29.373" r="3.378" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.053" cy="29.373" r="3.378" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.307" cy="40.122" r="3.378" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}