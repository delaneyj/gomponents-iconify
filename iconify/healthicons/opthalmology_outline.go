package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpthalmologyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M27.5 24.4a3.5 3.5 0 1 1-1.224-2.659a1.05 1.05 0 1 0 .382.382c.525.613.842 1.408.842 2.277Z"/><path fill-rule="evenodd" d="M24.037 16C30.3 16 38 24.4 38 24.4s-7.7 8.4-13.963 8.4C17.774 32.8 10 24.4 10 24.4S17.774 16 24.037 16Zm-11.202 8.4a40.797 40.797 0 0 1 3.055-2.615c.412-.318.839-.632 1.276-.936A7.668 7.668 0 0 0 16.3 24.4c0 1.28.313 2.488.866 3.55a32.419 32.419 0 0 1-1.276-.934a40.797 40.797 0 0 1-3.055-2.616Zm18.865 0c0-1.3-.322-2.526-.892-3.6c.457.317.903.647 1.334.98a40.278 40.278 0 0 1 3.036 2.62a40.278 40.278 0 0 1-3.036 2.62c-.43.333-.877.663-1.334.98c.57-1.074.892-2.3.892-3.6ZM24 30.1a5.7 5.7 0 1 0 0-11.4a5.7 5.7 0 0 0 0 11.4Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6 9a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v30a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9Zm2 0a1 1 0 0 1 1-1h30a1 1 0 0 1 1 1v30a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1V9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}