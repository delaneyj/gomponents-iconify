package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkullTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10v3.764a2 2 0 0 1-1.106 1.789L18 19v1a3 3 0 0 1-2.824 2.995L14.95 23c.022-.107.037-.218.044-.33L15 22.5V22a2 2 0 0 0-1.85-1.994L13 20h-2a2 2 0 0 0-1.995 1.85L9 22v.5c0 .171.017.339.05.5H9a3 3 0 0 1-3-3v-1l-2.894-1.447A2 2 0 0 1 2 15.763V12C2 6.477 6.477 2 12 2Zm0 2a8 8 0 0 0-7.996 7.75L4 12v3.764l4 2v1.591l.075-.084a3.993 3.993 0 0 1 2.723-1.266L11 18l2.073.001l.223.01a3.99 3.99 0 0 1 2.55 1.177l.154.167v-1.591l4-2V12a8 8 0 0 0-8-8Zm-4 7a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm8 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/>`),
		g.Group(children),
	)
}