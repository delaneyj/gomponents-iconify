package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhilipsSonicare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.009 7.45c4.032-5.678 16.077-3.455 16.077 9.25c0 13.738-6.492 26.8-10.559 26.8c-2.777 0-2.972-1.826-2.972-3.332c0-1.191.678-5.919-2.555-5.919s-2.555 4.728-2.555 5.919c0 1.506-.196 3.332-2.972 3.332c-4.067 0-10.56-13.062-10.56-26.8c0-12.704 12.046-14.928 16.078-9.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 11.589c5.496-6.522 11.402-2.919 12.44 1.846c2.063 9.463-9.692 17.48-12.44 17.48c-2.748 0-14.503-8.017-12.44-17.48C12.598 8.67 18.504 5.067 24 11.589Z"/>`),
		g.Group(children),
	)
}