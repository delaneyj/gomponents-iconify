package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.684 4.025a8 8 0 1 0 7.287 7.287a7.935 7.935 0 0 0-.603-2.439l1.5-1.502A9.935 9.935 0 0 1 22 11.997c0 5.522-4.477 10-10 10s-10-4.478-10-10c0-5.523 4.477-10 10-10a9.981 9.981 0 0 1 4.626 1.132l-1.501 1.5a7.941 7.941 0 0 0-2.44-.604Zm7.801-1.928L21.9 3.511l-9.193 9.193l-1.412.002l-.002-1.416l9.192-9.193Z"/>`),
		g.Group(children),
	)
}