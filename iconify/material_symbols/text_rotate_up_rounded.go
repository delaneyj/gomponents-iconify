package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotateUpRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 20q-.425 0-.713-.288T17 19V6.8l-.35.35q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7L17.3 3.7q.15-.15.338-.225t.387-.075q.2 0 .388.075t.337.225l2.025 2.025q.3.3.3.7t-.3.7q-.3.3-.7.288t-.725-.288L19 6.8V19q0 .425-.288.713T18 20Zm-5.225-3.35l-8.85-3.3q-.375-.15-.65-.537T3 12q0-.425.275-.813t.65-.537l8.85-3.3q.5-.2.862.063t.363.812q0 .25-.163.475t-.387.3l-2.25.75v4.45l2.25.8q.225.075.388.3t.162.5q0 .55-.363.8t-.862.05Zm-3.175-3v-3.3l-4.55 1.6v.1l4.55 1.6Z"/>`),
		g.Group(children),
	)
}