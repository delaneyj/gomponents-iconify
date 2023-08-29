package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnarchiveRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17.675q.425 0 .713-.288t.287-.712v-3.2l.9.9q.275.275.7.275t.7-.275q.275-.275.275-.7t-.275-.7L12.7 10.4q-.3-.3-.7-.3t-.7.3l-2.6 2.575q-.275.275-.275.7t.275.7q.275.275.7.275t.7-.275l.9-.9v3.2q0 .425.288.713t.712.287ZM5 21q-.825 0-1.413-.588T3 19V6.525q0-.35.113-.675t.337-.6L4.7 3.725q.275-.35.687-.538T6.25 3h11.5q.45 0 .863.188t.687.537l1.25 1.525q.225.275.338.6t.112.675V19q0 .825-.588 1.413T19 21H5Zm.4-15h13.2l-.85-1H6.25L5.4 6Z"/>`),
		g.Group(children),
	)
}