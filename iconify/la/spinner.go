package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spinner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3a3 3 0 1 0 .002 6.002A3 3 0 0 0 16 3zM8.937 6.438a2.502 2.502 0 0 0-2.5 2.5c0 1.378 1.122 2.5 2.5 2.5a2.5 2.5 0 0 0 0-5zm14.126 1.5c-.551 0-1 .449-1 1c0 .55.449 1 1 1c.55 0 1-.45 1-1c0-.551-.45-1-1-1zM6 13.75a2.25 2.25 0 1 0 .001 4.501A2.25 2.25 0 0 0 6 13.75zm20 1a1.25 1.25 0 1 0-.002 2.498A1.25 1.25 0 0 0 26 14.75zM8.937 21.063c-1.105 0-2 .894-2 2a1.999 1.999 0 1 0 4 0c0-1.106-.894-2-2-2zm14.126.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zM16 24.25c-.965 0-1.75.785-1.75 1.75s.785 1.75 1.75 1.75s1.75-.785 1.75-1.75s-.785-1.75-1.75-1.75z"/>`),
		g.Group(children),
	)
}