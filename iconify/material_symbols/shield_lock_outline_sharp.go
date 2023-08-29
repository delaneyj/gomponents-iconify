package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldLockOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-3.475-.875-5.738-3.988T4 11.1V5l8-3l8 3v6.1q0 3.8-2.263 6.913T12 22Zm0-2.1q2.6-.825 4.3-3.3t1.7-5.5V6.375l-6-2.25l-6 2.25V11.1q0 3.025 1.7 5.5t4.3 3.3Zm0-7.9Zm-3 4h6v-5h-1v-1q0-.825-.588-1.413T12 8q-.825 0-1.413.588T10 10v1H9v5Zm2-5v-1q0-.425.288-.713T12 9q.425 0 .713.288T13 10v1h-2Z"/>`),
		g.Group(children),
	)
}