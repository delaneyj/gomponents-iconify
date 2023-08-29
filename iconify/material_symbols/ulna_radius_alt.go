package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UlnaRadiusAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 12.775L7.5 4q1.875.925 3.475 2.538t2.4 3.362q.9-.95 1.988-1.525T18.4 7.3q1.15-.325 1.988-.625T22 6v9.65l-7.5 4.725q-.475.3-1.012.463T12.35 21q-.75 0-1.45-.288T9.625 19.9L2 12.775ZM12.75 17.5q.425 0 .713-.288t.287-.712q0-.05-.1-.425l3.55-2.225q.125.2.338.3t.462.1q.425 0 .712-.288T19 13.25q0-.35-.225-.625t-.575-.35q.025-.05.038-.125t.012-.15q0-.425-.288-.713T17.25 11q-.425 0-.713.288T16.25 12q0 .1.175.55l-3.55 2.2q-.125-.225-.363-.363T12 14.25q-.425 0-.713.288T11 15.25q0 .35.225.625t.575.35q-.025.05-.038.125t-.012.15q0 .425.288.713t.712.287Z"/>`),
		g.Group(children),
	)
}