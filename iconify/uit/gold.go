package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 11h9.1c.3 0 .5-.3.4-.6l-1-6c0-.2-.3-.4-.5-.4h-7c-.2 0-.5.2-.5.4l-1 6v.1c0 .3.2.5.5.5zm1.4-6h6.2l.8 5H8.1l.8-5zM22 13.4c0-.2-.2-.4-.5-.4h-7c-.2 0-.5.2-.5.4l-1 6v.1c0 .3.2.5.5.5h9.1c.3 0 .5-.3.4-.6l-1-6zM14.1 19l.8-5h6.2l.8 5h-7.8zm-4.6-6h-7c-.2 0-.5.2-.5.4l-1 6v.1c0 .3.2.5.5.5h9.1c.3 0 .5-.3.4-.6l-1-6c0-.2-.3-.4-.5-.4zm-7.4 6l.8-5h6.2l.8 5H2.1z"/>`),
		g.Group(children),
	)
}