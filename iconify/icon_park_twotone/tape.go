package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTape0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 12a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v24a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V12Z"/><path stroke-linecap="round" d="M14.27 10c-1.696 0-2.622 1.978-1.536 3.28l1.666 2a2 2 0 0 0 1.537.72h16.126a2 2 0 0 0 1.537-.72l1.666-2c1.086-1.302.16-3.28-1.536-3.28H14.27Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M6 10v0h36"/><path fill="#555" d="M33 31a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm-18 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTape0)"/>`),
		g.Group(children),
	)
}