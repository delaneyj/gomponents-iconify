package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Define(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.4 36.8c2.465 0 2.8-.885 2.8-3.8V14.7c0-2.075.227-3.8-2.7-3.8m-.9 0h11.6c3.1 0 5.6.4 7.5 1.3c2.3 1 4 2.6 5.2 4.7s1.8 4.4 1.8 7.1c0 11.383-7.445 12.8-14.6 12.8H12.5m9.23-25.9v22.648c0 3.142 1.275 3.252 2.467 3.252c2.354 0 4.148-.785 5.494-2.467c1.682-2.242 2.579-5.606 2.579-10.203c0-3.7-.56-6.727-1.794-8.97c-.897-1.793-2.13-2.915-3.588-3.587c-1.01-.449-2.691-.673-5.158-.673Z"/>`),
		g.Group(children),
	)
}