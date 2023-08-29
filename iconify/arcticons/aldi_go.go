package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AldiGo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.217 11.384L12.5 31.966m12.465-20.582l-7.717 20.582m12.464-20.582l-7.716 20.582m8.615-15.436h1.054m-2.983 5.145h4.889m-6.819 5.146H35.5M37.323 4.5H10.677a4 4 0 0 0-4 4v26.646a4 4 0 0 0 4 4h16.627V43.5l4.58-4.354h5.44a4 4 0 0 0 4-4V8.5a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}