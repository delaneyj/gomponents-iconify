package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestCamWallMountOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17.45q-.825 0-1.413-.588T2 15.45v-8q0-.825.588-1.412T4 5.45q.825 0 1.413.6T6 7.475q.875-1.375 2.313-2.2T11.5 4.45q1.3 0 2.488.488T16.1 6.35l4.3 4.3q.3.3.45.663t.15.762q0 .4-.15.763t-.45.662l-6.35 6.35q-.275.275-.638.425t-.762.15q-.4 0-.775-.15t-.65-.425L6.9 15.55q-.275-.275-.488-.55T6 14.425v1.025q0 .825-.588 1.413T4 17.45Zm3-6.525q0 .9.338 1.725t.987 1.475l4.325 4.325L19 12.075L14.675 7.75q-.65-.65-1.462-.988T11.5 6.425q-1.875 0-3.188 1.313T7 10.925Zm6 1.525Z"/>`),
		g.Group(children),
	)
}