package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Obelisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.25 12.25h-.37L11 2.87a1.21 1.21 0 0 0-.38-.78L8.87.35a1.25 1.25 0 0 0-1.74 0L5.34 2.09a1.21 1.21 0 0 0-.34.78l-.84 9.38h-.41A1.26 1.26 0 0 0 2.5 13.5V16h11v-2.5a1.26 1.26 0 0 0-1.25-1.25zM6.2 3L8 1.25L9.8 3l.83 9.27H5.37zm6.05 11.77h-8.5V13.5h8.5z"/><circle cx="8.02" cy="4.04" r=".62" fill="currentColor"/>`),
		g.Group(children),
	)
}