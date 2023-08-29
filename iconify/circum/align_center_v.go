package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.446 11.5h-.51V9.07a2.5 2.5 0 0 0-2.5-2.5h-2.43a2.5 2.5 0 0 0-2.5 2.5v2.43H11.5V6.58A2.5 2.5 0 0 0 9 4.08H6.566a2.5 2.5 0 0 0-2.5 2.5v4.92h-.52a.5.5 0 0 0 0 1h.52v4.92a2.5 2.5 0 0 0 2.5 2.5H9a2.5 2.5 0 0 0 2.5-2.5V12.5h1.01v2.43a2.5 2.5 0 0 0 2.5 2.5h2.43a2.5 2.5 0 0 0 2.5-2.5V12.5h.51a.5.5 0 0 0-.004-1ZM10.5 17.42a1.5 1.5 0 0 1-1.5 1.5H6.566a1.5 1.5 0 0 1-1.5-1.5V12.5H10.5Zm0-5.92H5.066V6.58a1.5 1.5 0 0 1 1.5-1.5H9a1.5 1.5 0 0 1 1.5 1.5Zm8.44 3.43a1.5 1.5 0 0 1-1.5 1.5h-2.43a1.5 1.5 0 0 1-1.5-1.5V12.5h5.43Zm0-3.43h-5.43V9.07a1.5 1.5 0 0 1 1.5-1.5h2.43a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		g.Group(children),
	)
}