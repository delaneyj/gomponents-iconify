package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceUnknown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q-.425 0-.738-.313t-.312-.737q0-.425.313-.737T12 14.9q.425 0 .738.313t.312.737q0 .425-.313.738T12 17Zm-.75-3.2q0-1.15.188-1.575T12.5 11.05q.35-.35.6-.688t.25-.762q0-.45-.338-.8T12 8.45q-.675 0-1.025.388t-.475.812L9.15 9.1q.3-.875 1.025-1.488T12 7q1.175 0 2.013.638t.837 1.912q0 .6-.3 1.125t-.75.975q-.75.75-.9 1.05t-.15 1.1h-1.5ZM7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Zm0-5h10V6H7v12Z"/>`),
		g.Group(children),
	)
}