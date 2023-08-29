package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotelClassRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.05 21.05q-.275.225-.587.025t-.188-.55L7.15 14.4l-4.875-3.5q-.3-.2-.187-.55T2.55 10H8.6l1.925-6.4q.05-.2.188-.275T11 3.25q.15 0 .288.075t.187.275L13.4 10h6.05q.35 0 .463.35t-.188.55l-4.875 3.5l1.875 6.125q.125.35-.188.55t-.587-.025L11 17.3l-4.95 3.75Zm14.05 0l-1.475-1.125L17.15 15.2l3.1-2.2h1.875q.35 0 .475.35t-.175.55L19.5 16l1.4 4.5q.125.35-.2.563t-.6-.013ZM14.9 8l-.75-2.55l.55-1.85q.05-.2.188-.288t.287-.087q.15 0 .288.1t.187.275L17 8h-2.1Z"/>`),
		g.Group(children),
	)
}