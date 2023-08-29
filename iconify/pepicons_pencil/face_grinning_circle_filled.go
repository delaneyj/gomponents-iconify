package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceGrinningCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilFaceGrinningCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M12.75 20.5a7.75 7.75 0 1 0 0-15.5a7.75 7.75 0 0 0 0 15.5Zm0 1a8.75 8.75 0 1 0 0-17.5a8.75 8.75 0 0 0 0 17.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.856 14.692a.5.5 0 0 1 .394-.192h7.5a.5.5 0 0 1 .485.621L16.75 15l.485.121v.002l-.001.004l-.003.01l-.01.038a7.562 7.562 0 0 1-.193.587c-.14.372-.359.873-.674 1.378c-.623.996-1.692 2.11-3.354 2.11c-1.662 0-2.731-1.114-3.354-2.11a7.806 7.806 0 0 1-.867-1.965l-.01-.038l-.003-.01v-.004c0-.001-.001-.002.484-.123l-.485.122a.502.502 0 0 1 .09-.43Zm1.087.808c.124.318.305.716.551 1.11c.55.88 1.355 1.64 2.506 1.64c1.15 0 1.956-.76 2.506-1.64c.246-.394.427-.792.551-1.11H9.943Z" clip-rule="evenodd"/><path d="M11.5 10.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm5.5 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilFaceGrinningCircleFilled0)"/></g>`),
		g.Group(children),
	)
}