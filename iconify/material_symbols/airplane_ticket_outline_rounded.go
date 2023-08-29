package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneTicketOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.8 12.9l-2.4.6l-1.075-.8q-.075-.05-.4-.1l-.125.05q-.225.05-.325.263t.025.412l1.15 2q.1.15.25.213t.325.012l8.525-2.25q.375-.1.563-.463t.087-.737q-.1-.375-.437-.562t-.713-.088l-2.45.65l-3.725-3.5q-.125-.125-.3-.162t-.35.012l-.125.025q-.35.075-.488.4t.038.625l1.95 3.4ZM4 20q-.825 0-1.413-.588T2 18v-3.375q0-.275.175-.475t.45-.25q.6-.2.988-.725T4 12q0-.65-.388-1.175t-.987-.725q-.275-.05-.45-.25T2 9.375V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V6H4v2.55q.925.55 1.463 1.463T6 12q0 1.075-.537 1.988T4 15.45V18Zm8-6Z"/>`),
		g.Group(children),
	)
}