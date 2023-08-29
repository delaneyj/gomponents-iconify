package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCabinet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFileCabinet0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M42 4H6v10h36V4Zm0 15H6v10h36V19Zm0 15H6v10h36V34Z"/><path stroke="#000" stroke-linecap="round" d="M21 9h6m-6 15h6m-6 15h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFileCabinet0)"/>`),
		g.Group(children),
	)
}