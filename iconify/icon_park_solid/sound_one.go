package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSoundOne0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-width="4" d="M9 7a3 3 0 0 1 3-3h24a3 3 0 0 1 3 3v34a3 3 0 0 1-3 3H12a3 3 0 0 1-3-3V7Z"/><path fill="#000" stroke="#000" stroke-linejoin="round" stroke-width="4" d="M24 29a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><rect width="4" height="4" x="30" y="8" fill="#000" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M27 36h2m-10 0h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSoundOne0)"/>`),
		g.Group(children),
	)
}