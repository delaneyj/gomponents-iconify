package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-width="2" d="m10.536 32.364l.707-.707l-.707.707l5.1 5.1a2 2 0 1 0 2.828-2.828l-5.1-5.1a2 2 0 0 0-2.828 0l.707.707l-.707-.707a2 2 0 0 0 0 2.828Zm28.147-3l-.924.381l.924-.38a2 2 0 0 0-2.61-1.088l-6.67 2.748a2 2 0 1 0 1.525 3.698l6.668-2.748a2 2 0 0 0 1.087-2.61ZM31.5 14.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0Zm-5.206-3.427a2 2 0 0 0-3.986-.333l-.602 7.187a2 2 0 1 0 3.986.334l.602-7.188ZM22 33.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0Zm19-2a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0Z"/>`),
		g.Group(children),
	)
}