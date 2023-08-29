package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessDownFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 8a4 4 0 1 1-3.995 4.2L8 12l.005-.2A4 4 0 0 1 12 8zm0-4a1 1 0 0 1 .993.883L13 5.01a1 1 0 0 1-1.993.117L11 5a1 1 0 0 1 1-1zm5 2a1 1 0 0 1 .993.883L18 7.01a1 1 0 0 1-1.993.117L16 7a1 1 0 0 1 1-1zm2 5a1 1 0 0 1 .993.883l.007.127a1 1 0 0 1-1.993.117L18 12a1 1 0 0 1 1-1zm-2 5a1 1 0 0 1 .993.883l.007.127a1 1 0 0 1-1.993.117L16 17a1 1 0 0 1 1-1zm-5 2a1 1 0 0 1 .993.883l.007.127a1 1 0 0 1-1.993.117L11 19a1 1 0 0 1 1-1zm-5-2a1 1 0 0 1 .993.883L8 17.01a1 1 0 0 1-1.993.117L6 17a1 1 0 0 1 1-1zm-2-5a1 1 0 0 1 .993.883L6 12.01a1 1 0 0 1-1.993.117L4 12a1 1 0 0 1 1-1zm2-5a1 1 0 0 1 .993.883L8 7.01a1 1 0 0 1-1.993.117L6 7a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}