package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ring(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#ECF0F1" d="M50-.027c27.613 0 50 22.386 50 50c0 27.613-22.387 50-50 50c-27.614 0-50-22.387-50-50c0-27.614 22.386-50 50-50z"/><defs><circle id="flatUiRing0" cx="50" cy="49.973" r="50"/></defs><clipPath id="flatUiRing1"><use href="#flatUiRing0"/></clipPath><g clip-path="url(#flatUiRing1)"><path fill="#526476" d="M-1 71.973h80v49H-1z"/><path fill="#F1C40E" d="M29.5 56.973h41a2.499 2.499 0 1 1 0 5h-41a2.5 2.5 0 1 1 0-5z"/><path fill="#F39C12" d="M30 52.973h40v4H30v-4z"/><path fill="#E79410" d="M30 52.973h20v4H30v-4z"/><path fill="#F1C40E" d="M47 33.973h6c10.493 0 19 8.507 19 19H28c0-10.493 8.507-19 19-19z"/><path fill="#F39C12" d="M54 29.473a1.5 1.5 0 0 0-1.5-1.5h-5a1.5 1.5 0 0 0 0 3H47a2 2 0 0 1 2 2v1h2v-1a2 2 0 0 1 2-2h-.5a1.5 1.5 0 0 0 1.5-1.5z"/><path fill="#E79410" d="M47.5 27.973a1.5 1.5 0 0 0 0 3H47a2 2 0 0 1 2 2v1h1v-6h-2.5z"/><path fill="#E5BA0D" d="M50 33.973h-3c-10.493 0-19 8.507-19 19h22v-19zm-20.5 23a2.499 2.499 0 1 0 0 5H50v-5H29.5z"/><path fill="#455564" d="M-1 72.973v48h37l43-43v-5z"/><path fill="#2C3E50" d="M-1 61.973h85v12H-1z"/></g>`),
		g.Group(children),
	)
}