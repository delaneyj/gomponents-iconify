package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filevideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m5.65 11.94l4.14-2.56c.28-.17.28-.59 0-.76L5.65 6.06c-.29-.18-.65.04-.65.38v5.11c0 .34.36.56.65.38Z"/><path fill="currentColor" d="M12.5 16h-10c-.83 0-1.5-.67-1.5-1.5v-13C1 .67 1.67 0 2.5 0h7.09c.4 0 .78.16 1.06.44l2.91 2.91c.28.28.44.66.44 1.06V14.5c0 .83-.67 1.5-1.5 1.5ZM2.5 1c-.28 0-.5.22-.5.5v13c0 .28.22.5.5.5h10c.28 0 .5-.22.5-.5V4.41a.47.47 0 0 0-.15-.35L9.94 1.15A.51.51 0 0 0 9.59 1H2.5Z"/><path fill="currentColor" d="M13.38 5h-2.91C9.66 5 9 4.34 9 3.53V.62c0-.28.22-.5.5-.5s.5.22.5.5v2.91c0 .26.21.47.47.47h2.91c.28 0 .5.22.5.5s-.22.5-.5.5Z"/>`),
		g.Group(children),
	)
}