package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DangerTriangleLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M5.312 10.761C8.23 5.587 9.689 3 12 3c2.31 0 3.77 2.587 6.688 7.761l.364.645c2.425 4.3 3.638 6.45 2.542 8.022S17.786 21 12.364 21h-.728c-5.422 0-8.134 0-9.23-1.572s.117-3.722 2.542-8.022l.364-.645Z" opacity=".5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 8v5"/><circle cx="12" cy="16" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}