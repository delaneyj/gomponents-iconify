package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Injection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSInjection0"><g fill="none"><path fill="#fff" fill-rule="evenodd" d="M38.168 22.262L19.077 41.354L6.349 28.626L25.44 9.534" clip-rule="evenodd"/><path stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M38.168 22.262L19.077 41.354L6.349 28.626L25.44 9.534"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="m21.905 5.999l19.8 19.799"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="m14.834 28.626l4.243 4.243"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="m6.35 41.353l6.363-6.363m19.092-19.092l3.534-3.535"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSInjection0)"/>`),
		g.Group(children),
	)
}