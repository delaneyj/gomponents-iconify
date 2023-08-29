package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Appwash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="39.7" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.5 4.5h-27c-1.1 0-2 .9-2 2v35c0 1.1.9 2 2 2h27c1.1 0 2-.9 2-2v-35c0-1.1-.9-2-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.4 8.1H13.6c-.8 0-1.5.6-1.5 1.4v24.9c0 .8.7 1.4 1.5 1.4h20.7c.8 0 1.5-.6 1.5-1.4V9.6c.1-.8-.6-1.5-1.4-1.5Z"/><circle cx="24" cy="24" r="8.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="5.2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.8 24.2c5.9-2.3 4.4 1.2 10.2.6"/><circle cx="30.7" cy="10.9" r=".75" fill="currentColor"/><circle cx="33" cy="10.9" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}