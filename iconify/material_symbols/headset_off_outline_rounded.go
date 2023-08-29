package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.1 18.15l-2-2V14h-2v.15l-1.475-1.475q.25-.3.625-.487T17.1 12h2v-1q0-2.95-2.05-4.975T12.1 4q-1.1 0-2.087.313T8.2 5.2L6.75 3.8q1.125-.875 2.488-1.338T12.1 2q1.85 0 3.488.7t2.862 1.925q1.225 1.225 1.938 2.863T21.1 11v7.15Zm-8 4.85q-.425 0-.712-.288T12.1 22q0-.425.288-.713T13.1 21h5l-1-1q-.825 0-1.413-.588T15.1 18L5.575 8.475q-.225.6-.35 1.225T5.1 11v1h2q.825 0 1.413.588T9.1 14v4q0 .825-.588 1.413T7.1 20h-2q-.825 0-1.412-.588T3.1 18v-7q0-1.125.238-2.113t.737-1.937L1.4 4.275q-.3-.3-.287-.7t.312-.7q.3-.275.7-.288t.7.288l18.375 18.4q.275.275.275.688t-.275.712q-.3.3-.687.325t-.588-.175q-.15.125-.35.15T19.1 23h-6Zm-8-5h2v-4h-2v4Zm0 0h2h-2Zm14-1.85Z"/>`),
		g.Group(children),
	)
}