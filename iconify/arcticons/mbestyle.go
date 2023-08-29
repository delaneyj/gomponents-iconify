package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mbestyle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.658 37.556h7.365v3.262A3.682 3.682 0 0 1 22.34 44.5h0a3.682 3.682 0 0 1-3.682-3.682v-3.262h0Zm6.818-32.271h-6.271c-8.502 0-12 3.582-12 8v21.271a3 3 0 0 0 3 3h24.27a3 3 0 0 0 3-3V13.285c0-4.418-3.368-8-12-8Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.476 37.556h3.32a3 3 0 0 0 3-3V13.285c0-4.418-3.37-8-12-8h-3.32"/><circle cx="17.521" cy="15.776" r="1.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="27.16" cy="15.776" r="1.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.157 19.307a2.933 2.933 0 0 1-5.633 0"/><circle cx="13.927" cy="22.042" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.66" cy="22.042" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}