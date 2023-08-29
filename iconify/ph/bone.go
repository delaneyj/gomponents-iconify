package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M231.67 60.89a35.82 35.82 0 0 0-23.82-12.74a36 36 0 1 0-66.37 22.92a.25.25 0 0 1 0 .08l-70.31 70.36h-.1a36 36 0 1 0-22.92 66.37a36 36 0 1 0 66.37-22.92a.54.54 0 0 1 0-.08l70.35-70.36h.1a36 36 0 0 0 46.74-53.63ZM219.1 97.16a20 20 0 0 1-25.67 3.8a16 16 0 0 0-19.88 2.19l-70.4 70.4a16 16 0 0 0-2.15 19.88a20 20 0 1 1-36.75 7.5a8 8 0 0 0-7.91-9.24a8.5 8.5 0 0 0-1.23.1A20 20 0 1 1 62.57 155a16 16 0 0 0 19.88-2.19l70.4-70.4A16 16 0 0 0 155 62.57a20 20 0 1 1 36.75-7.5a8 8 0 0 0 9.14 9.14a20 20 0 0 1 18.17 33Z"/>`),
		g.Group(children),
	)
}