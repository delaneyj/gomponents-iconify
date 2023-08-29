package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Server(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.437 11H3.563a1.5 1.5 0 0 1-1.5-1.5V5.565a1.5 1.5 0 0 1 1.5-1.5h16.874a1.5 1.5 0 0 1 1.5 1.5v3.93a1.5 1.5 0 0 1-1.5 1.505ZM3.563 5.065a.5.5 0 0 0-.5.5v3.93a.5.5 0 0 0 .5.5h16.874a.5.5 0 0 0 .5-.5v-3.93a.5.5 0 0 0-.5-.5Zm16.874 14.87H3.563a1.5 1.5 0 0 1-1.5-1.5v-3.93a1.5 1.5 0 0 1 1.5-1.5h16.874a1.5 1.5 0 0 1 1.5 1.5v3.93a1.5 1.5 0 0 1-1.5 1.5Zm-16.874-5.93a.5.5 0 0 0-.5.5v3.93a.5.5 0 0 0 .5.5h16.874a.5.5 0 0 0 .5-.5v-3.93a.5.5 0 0 0-.5-.5Z"/><circle cx="5.563" cy="7.53" r=".5" fill="currentColor"/><circle cx="7.563" cy="7.53" r=".5" fill="currentColor"/><path fill="currentColor" d="M13.452 8.03a.5.5 0 0 1 0-1h5a.5.5 0 0 1 0 1Z"/><circle cx="5.563" cy="16.47" r=".5" fill="currentColor"/><circle cx="7.563" cy="16.47" r=".5" fill="currentColor"/><path fill="currentColor" d="M13.452 16.97a.5.5 0 0 1 0-1h5a.5.5 0 0 1 0 1Z"/>`),
		g.Group(children),
	)
}