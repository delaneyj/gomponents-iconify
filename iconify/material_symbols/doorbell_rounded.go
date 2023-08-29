package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorbellRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17.5q.425 0 .713-.288T13 16.5h-2q0 .425.288.713T12 17.5ZM8.5 16h7q.2 0 .35-.15t.15-.35q0-.2-.15-.35T15.5 15H15v-2.35q0-1.1-.6-2T12.75 9.5v-.25q0-.325-.213-.537T12 8.5q-.325 0-.537.213t-.213.537v.25q-1.05.25-1.65 1.15t-.6 2V15h-.5q-.2 0-.35.15T8 15.5q0 .2.15.35t.35.15ZM6 21q-.825 0-1.413-.588T4 19v-9q0-.475.213-.9t.587-.7l6-4.5q.275-.2.575-.3T12 3.5q.325 0 .625.1t.575.3l6 4.5q.375.275.588.7T20 10v9q0 .825-.588 1.413T18 21H6Z"/>`),
		g.Group(children),
	)
}