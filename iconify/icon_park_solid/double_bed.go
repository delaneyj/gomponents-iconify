package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleBed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDoubleBed0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 12a3 3 0 0 1 3-3h26a3 3 0 0 1 3 3v11H8V12ZM6 35v4m36-4v4"/><path fill="#fff" d="M20 18h-6a3 3 0 0 0-3 3v2h12v-2a3 3 0 0 0-3-3Zm14 0h-6a3 3 0 0 0-3 3v2h12v-2a3 3 0 0 0-3-3Z"/><path d="M4 26a3 3 0 0 1 3-3h34a3 3 0 0 1 3 3v9H4v-9Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDoubleBed0)"/>`),
		g.Group(children),
	)
}