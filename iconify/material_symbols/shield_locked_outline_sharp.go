package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldLockedOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12Zm0 10q-3.475-.875-5.738-3.988T4 11.1V5l8-3l8 3v6.1q0 .25-.013.5t-.037.5q-.225-.05-.463-.075T19 12q-.275 0-.525.025t-.525.075q.025-.25.038-.487T18 11.1V6.375l-6-2.25l-6 2.25V11.1q0 3.025 1.7 5.5t4.3 3.3q.525-.175 1.025-.425T14 18.9v2.35q-.475.25-.975.438T12 22Zm4 0v-5h1v-1q0-.825.588-1.413T19 14q.825 0 1.413.588T21 16v1h1v5h-6Zm2-5h2v-1q0-.425-.288-.713T19 15q-.425 0-.713.288T18 16v1Z"/>`),
		g.Group(children),
	)
}