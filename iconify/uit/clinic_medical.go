package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClinicMedical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 12v1.5H10c-.3 0-.5.2-.5.5s.2.5.5.5h1.5V16c0 .3.2.5.5.5s.5-.2.5-.5v-1.5H14c.3 0 .5-.2.5-.5s-.2-.5-.5-.5h-1.5V12c0-.3-.2-.5-.5-.5s-.5.2-.5.5zm10.3-1.4l-9.5-8.4c-.2-.2-.5-.2-.7 0l-9.5 8.4c-.2.2-.2.5 0 .7c.2.2.5.2.7 0l1.2-1v11.2c0 .3.2.5.5.5h15c.3 0 .5-.2.5-.5V10.3l1.2 1c.1.1.2.1.3.1c.1 0 .3-.1.4-.2c.2-.1.1-.5-.1-.6zM19 21H5V9.4l7-6.2l7 6.2V21z"/>`),
		g.Group(children),
	)
}