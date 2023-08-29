package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoStableOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.025 12Zm5.2 5.025l1.8-6.725q.125-.4-.087-.75t-.613-.475L8.05 6.275q-.4-.125-.75.088t-.475.612l-1.8 6.725q-.125.4.088.75t.612.475L16 17.725q.4.125.75-.088t.475-.612ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}