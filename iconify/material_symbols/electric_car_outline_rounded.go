package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricCarOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 15v.5q0 .625-.438 1.063T4.5 17q-.625 0-1.063-.438T3 15.5V8.35q0-.175.025-.35t.075-.325L4.975 2.35q.2-.6.725-.975T6.875 1h10.25q.65 0 1.175.375t.725.975L20.9 7.675q.05.15.075.325t.025.35v7.15q0 .625-.438 1.063T19.5 17q-.625 0-1.063-.438T18 15.5V15H6Zm-.2-9h12.4l-1.05-3H6.85L5.8 6ZM5 8v5v-5Zm2.5 4q.625 0 1.063-.438T9 10.5q0-.625-.438-1.063T7.5 9q-.625 0-1.063.438T6 10.5q0 .625.438 1.063T7.5 12Zm9 0q.625 0 1.063-.438T18 10.5q0-.625-.438-1.063T16.5 9q-.625 0-1.063.438T15 10.5q0 .625.438 1.063T16.5 12ZM13 21v1.2q0 .275-.238.425t-.487.025L7.95 20.475q-.175-.1-.138-.288T8.05 20H11v-1.2q0-.275.238-.425t.487-.025l4.325 2.175q.175.1.138.288T15.95 21H13Zm-8-8h14V8H5v5Z"/>`),
		g.Group(children),
	)
}