package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18.3V5.7c.6-.3 1-1 1-1.7c0-1.1-.9-2-2-2c-.7 0-1.4.4-1.7 1H5.7c-.3-.6-1-1-1.7-1c-1.1 0-2 .9-2 2c0 .7.4 1.4 1 1.7v12.6c-.6.3-1 1-1 1.7c0 1.1.9 2 2 2c.7 0 1.4-.4 1.7-1h12.6c.3.6 1 1 1.7 1c1.1 0 2-.9 2-2c0-.7-.4-1.4-1-1.7zm-2 0c-.3.2-.5.4-.7.7H5.7c-.2-.3-.4-.5-.7-.7V5.7c.3-.2.5-.4.7-.7h12.6c.2.3.4.5.7.7v12.6zM14 9V8c0-.6-.4-1-1-1H8c-.6 0-1 .4-1 1v5c0 .6.4 1 1 1h1v-3c0-1.1.9-2 2-2h3zm2 1h-5c-.6 0-1 .4-1 1v5c0 .6.4 1 1 1h5c.6 0 1-.4 1-1v-5c0-.6-.4-1-1-1z"/>`),
		g.Group(children),
	)
}