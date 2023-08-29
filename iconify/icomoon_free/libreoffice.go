package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libreoffice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.354.354A1.394 1.394 0 0 0 7.5 0h-6c-.275 0-.5.225-.5.5v15c0 .275.225.5.5.5h12c.275 0 .5-.225.5-.5v-9c0-.275-.159-.659-.354-.854L8.353.353zM13 15H2V1h5.487a.545.545 0 0 1 .169.07l5.274 5.274a.545.545 0 0 1 .07.169V15zm.5-15h-3c-.275 0-.341.159-.146.354l3.293 3.293c.194.194.354.129.354-.146v-3c0-.275-.225-.5-.5-.5z"/>`),
		g.Group(children),
	)
}