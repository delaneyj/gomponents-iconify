package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PermDeviceInformationSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 16.5V11h2v5.5h-2ZM12 9q-.425 0-.713-.288T11 8q0-.425.288-.713T12 7q.425 0 .713.288T13 8q0 .425-.288.713T12 9ZM5 23V1h14v22H5Zm2-5h10V6H7v12Z"/>`),
		g.Group(children),
	)
}