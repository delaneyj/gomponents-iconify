package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Massive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.843 7.558L20.785 5.5L18.12 8.166l-2.652-2.653L11.6 9.38L9.676 7.456l-2.14 2.14l1.963 1.964l-3.937 3.936l2.598 2.585l-2.64 2.649l2.095 2.094L22.843 7.558Zm2.386 32.866l2.076 2.076l2.675-2.675l2.607 2.607l3.891-3.89l1.95 1.948l2.01-2.01l-2.02-2.019l3.968-3.968l-2.58-2.58l2.674-2.674l-2.05-2.05l-15.201 15.235ZM15.182 15.237L32.79 32.846"/>`),
		g.Group(children),
	)
}