package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freeagent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.296 31.453H24.567m-9.341 10.949L28.479 9.475L41.731 42.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.479 9.475c1.422-3.535 3.878-5.178 8.333-2.98M6.269 7.697c4.602 3.052 9.35.645 14.1.645c5.158 0 4.308 7.093 3.839 11.783c-1.524-2.902-3.518-3.752-5.306-3.4c-1.788.352-5.598 1.524-7.152.205s-4.778-5.104-5.48-9.233Z"/>`),
		g.Group(children),
	)
}