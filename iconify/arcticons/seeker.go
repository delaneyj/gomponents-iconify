package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seeker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.336 28.58s4.168 3.371 5.3 3.666m.01-8.302a13.516 13.516 0 0 1 24.267-7.821l6.565.353c1.068.057 1.694.85 1.09 1.67c-.329.328-4.911 5.498-4.911 5.498c-.916 1.142-6.094 6.174 4.225 14.847"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.637 32.246a57.31 57.31 0 0 1 .015-8.475m11.418-4.729a2.754 2.754 0 1 1 5.396-1.102"/>`),
		g.Group(children),
	)
}