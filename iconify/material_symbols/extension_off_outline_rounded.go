package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExtensionOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 17.15l-2-2V13h2q.2 0 .35-.15t.15-.35q0-.2-.15-.35T20 12h-2V6h-6V4q0-.2-.15-.35t-.35-.15q-.2 0-.35.15T11 4v2H8.85l-2-2H9q0-1.05.725-1.775T11.5 1.5q1.05 0 1.775.725T14 4h4q.825 0 1.413.588T20 6v4q1.05 0 1.775.725T22.5 12.5q0 1.05-.725 1.775T20 15v2.15Zm-6.55-6.55Zm5.625 11.3l-17-16.975q-.3-.3-.3-.712t.3-.713q.3-.3.713-.3t.712.3l17 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3Zm-8.5-8.475ZM5 21q-.825 0-1.413-.588T3 19v-3.8q1.2 0 2.1-.762T6 12.5q0-1.175-.9-1.937T3 9.8V6q0-.425.15-.788t.425-.637L5 6v2.2q1.35.5 2.175 1.675T8 12.5q0 1.425-.825 2.6T5 16.8V19h2.2q.525-1.35 1.7-2.175T11.5 16q1.425 0 2.6.825T15.8 19H18l1.425 1.425q-.275.275-.638.425T18 21h-3.8q0-1.2-.762-2.1T11.5 18q-1.175 0-1.938.9T8.8 21H5Z"/>`),
		g.Group(children),
	)
}