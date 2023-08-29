package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretLeftTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.927 18.839c.808.707 2.073.133 2.073-.94V6.101c0-1.074-1.265-1.648-2.073-.94l-6.31 5.521a1.75 1.75 0 0 0 0 2.634l6.31 5.522Zm.573-1.492l-5.896-5.159a.25.25 0 0 1 0-.376L13.5 6.653v10.694Z"/>`),
		g.Group(children),
	)
}