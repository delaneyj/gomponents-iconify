package bytesize

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 2c-4 0-4 4-4 4v10s0 4 4 4s4-4 4-4V6s0-4-4-4ZM8 17s0 7 8 7s8-7 8-7M13 29h6m-3-5v5"/>`),
		g.Group(children),
	)
}