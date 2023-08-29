package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceUnknownRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q-.425 0-.738-.313t-.312-.737q0-.425.313-.737T12 14.9q.425 0 .738.313t.312.737q0 .425-.313.738T12 17ZM9.725 9.325q-.25-.125-.325-.413t.1-.537q.425-.65 1.088-1.012T12 7q1.2 0 2.025.688t.825 1.862q0 .625-.313 1.15t-.737.95q-.325.325-.638.65t-.387.775q-.05.3-.263.513T12 13.8q-.3 0-.513-.2t-.212-.5q0-.625.375-1.125t.85-.925q.325-.325.588-.675t.262-.775q0-.5-.413-.825T12 8.45q-.35 0-.688.163t-.537.462q-.2.275-.488.338t-.562-.088ZM7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Zm0-5h10V6H7v12Z"/>`),
		g.Group(children),
	)
}