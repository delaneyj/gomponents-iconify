package geo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurfSquareGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<circle cx="88.472" cy="11.528" r="5.824" fill="currentColor"/><circle cx="88.472" cy="88.472" r="5.824" fill="currentColor"/><circle cx="50" cy="50" r="5.824" fill="currentColor"/><circle cx="11.528" cy="88.472" r="5.824" fill="currentColor"/><circle cx="11.528" cy="11.528" r="5.824" fill="currentColor"/><path fill="currentColor" d="M48 40.382V12H21.328c-.241 5.04-4.288 9.087-9.328 9.328V48h28.382A9.851 9.851 0 0 1 48 40.382zM59.618 48H88V21.328c-5.04-.241-9.087-4.288-9.329-9.328H52v28.382A9.851 9.851 0 0 1 59.618 48zm-19.236 4H12v26.671c5.04.241 9.087 4.289 9.328 9.329H48V59.618A9.851 9.851 0 0 1 40.382 52zM88 78.671V52H59.618A9.851 9.851 0 0 1 52 59.618V88h26.671c.241-5.04 4.289-9.088 9.329-9.329z"/>`),
		g.Group(children),
	)
}