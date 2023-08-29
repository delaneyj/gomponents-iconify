package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apoteket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.418 13.217c-5.932-5.307-19.325-3.805-24.712-.869s-6.047 5.993-6.19 9.878C4.3 28.019 6.288 35.51 12.633 37.602c8.474 2.794 17.934-3.545 17.934-3.545l.125 3.974l12.596-.312l.212-1.668l-4.7-.549s2.55-16.978-3.382-22.285Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.537 15.704c.032-2.761 15.604-3.769 15.631-.046L30.13 29.69c-2.961 5.307-15.326 6.484-15.473 2.872Z"/>`),
		g.Group(children),
	)
}