package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenDoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSOpenDoor0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M20 4v40l22-6V10L20 4Z"/><path stroke="#fff" stroke-linecap="round" d="M6 10h14v28H6V10Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="M28 22v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSOpenDoor0)"/>`),
		g.Group(children),
	)
}