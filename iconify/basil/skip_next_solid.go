package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipNextSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.75 7a.75.75 0 0 0-1.5 0v10a.75.75 0 0 0 1.5 0V7Zm-4.296 3.945c.69.534.69 1.576 0 2.11a25.51 25.51 0 0 1-7.073 3.863l-.466.166c-.87.308-1.79-.28-1.907-1.178a30.314 30.314 0 0 1 0-7.812c.118-.898 1.037-1.486 1.907-1.177l.466.165a25.511 25.511 0 0 1 7.073 3.863Z"/>`),
		g.Group(children),
	)
}