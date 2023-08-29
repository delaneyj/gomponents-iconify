package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FinanceChipOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.75 16h.725v-.8q.7-.1 1.188-.575t.487-1.225q0-.65-.5-1.088T12.5 11.6V9.75q.25.075.413.25t.237.425l.9-.375q-.175-.525-.6-.837T12.5 8.8V8h-.75v.775q-.7.075-1.188.513t-.487 1.162q0 .675.513 1.125t1.162.725v1.975q-.4-.125-.675-.425t-.375-.7l-.875.375q.2.7.7 1.15t1.225.55V16Zm.75-1.75V12.6q.275.125.488.3t.212.525q0 .4-.2.563t-.5.262Zm-.75-2.975q-.275-.125-.5-.3t-.225-.525q0-.35.225-.513t.5-.212v1.55ZM8 19q-2.925 0-4.963-2.038T1 12q0-2.925 2.038-4.963T8 5h8q2.925 0 4.963 2.038T23 12q0 2.925-2.038 4.963T16 19H8Zm0-2h8q2.075 0 3.538-1.463T21 12q0-2.075-1.463-3.538T16 7H8Q5.925 7 4.462 8.463T3 12q0 2.075 1.463 3.538T8 17Zm4-5Z"/>`),
		g.Group(children),
	)
}