package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WheelBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><g opacity=".5"><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Zm-10 6a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm3-6a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z" clip-rule="evenodd"/></g><path d="m9.672 17.532l1.525-2.64a2.997 2.997 0 0 1-1.299-.751l-1.524 2.64c.397.301.832.555 1.298.75ZM6.046 12.75h3.048a3.006 3.006 0 0 1 0-1.5H6.047a6.059 6.059 0 0 0 0 1.5Zm2.328-5.53l1.524 2.64a2.987 2.987 0 0 1 1.299-.751l-1.525-2.64a5.978 5.978 0 0 0-1.298.75Zm4.429 1.889l1.525-2.64c.465.195.901.449 1.298.75l-1.524 2.64a2.996 2.996 0 0 0-1.299-.75Zm2.102 3.641a3.004 3.004 0 0 0 0-1.5h3.049a6.069 6.069 0 0 1 0 1.5h-3.049Zm-2.102 2.141c.497-.137.942-.4 1.299-.75l1.524 2.64a5.993 5.993 0 0 1-1.298.75l-1.525-2.64Z"/></g>`),
		g.Group(children),
	)
}