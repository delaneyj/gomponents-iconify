package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlaskAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 15a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 2a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/><path fill="currentColor" fill-rule="evenodd" d="M14.495 3.25H8.5a.75.75 0 0 0 0 1.5h.26v5.061c-4.452 2.104-5.095 8.241-1.094 11.209l.295.218a6.78 6.78 0 0 0 8.078 0l.295-.218c4.001-2.968 3.358-9.104-1.094-11.209V4.75h.26a.75.75 0 0 0 0-1.5h-1.005Zm-5.64 16.784a5.28 5.28 0 0 0 6.29 0l.295-.22c2.016-1.494 2.504-4.03 1.644-6.064H6.915c-.86 2.033-.37 4.57 1.645 6.065l.295.218Zm-.958-7.784h8.206a5.07 5.07 0 0 0-1.676-1.16a1.138 1.138 0 0 1-.687-1.045V4.75h-3.48v5.295c0 .454-.27.865-.688 1.045a5.07 5.07 0 0 0-1.675 1.16Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}