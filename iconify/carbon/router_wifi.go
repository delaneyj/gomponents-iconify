package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RouterWifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="16" cy="13.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="M19.536 10.465a5 5 0 0 0-7.072 0L11.05 9.05a7 7 0 0 1 9.9 0Z"/><path fill="currentColor" d="M23.071 6.929a10 10 0 0 0-14.142 0L7.515 5.514a12.001 12.001 0 0 1 16.97 0zM21 25l-5 5l-5-5l1.409-1.419L15 26.153V19h2v7.206l2.591-2.625L21 25zm3-14l-5 5l5 5l1.419-1.409L22.847 17H30v-2h-7.206l2.625-2.591L24 11zM8 11l5 5l-5 5l-1.419-1.409L9.153 17H2v-2h7.206l-2.625-2.591L8 11z"/>`),
		g.Group(children),
	)
}