package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOffLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.587 21H3a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h.586L1.395 2.808l1.414-1.414l19.799 19.798l-1.415 1.415L19.587 21Zm-14-14H4v12h13.586l-2.18-2.18a5.5 5.5 0 0 1-7.725-7.725L5.587 7Zm3.524 3.525a3.5 3.5 0 0 0 4.865 4.865l-4.865-4.865Zm12.89 7.261l-2-2V7h-3.828l-2-2H9.83l-.308.307l-1.414-1.414L9.001 3h6l2 2h4a1 1 0 0 1 1 1v11.786ZM11.264 7.049a5.5 5.5 0 0 1 6.188 6.188l-2.338-2.338a3.516 3.516 0 0 0-1.512-1.512l-2.338-2.338Z"/>`),
		g.Group(children),
	)
}