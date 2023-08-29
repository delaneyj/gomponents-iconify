package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nekome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.574 24.579c0 .02-4.863-1.514-4.863-1.514m4.863 4.488H4.5m5.074 3.157L4.73 32.134m9.197-7.123s5.04-3.986 8.9.507"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 16.331c4.1 0 8.795-2.189 12.376-3.448s2.971.64 2.54 1.61a20.046 20.046 0 0 0-.535 4.345l.045 15.299a1.395 1.395 0 0 1-1.398 1.398H10.972a1.393 1.393 0 0 1-1.398-1.398l.045-15.299a20.046 20.046 0 0 0-.534-4.345c-.432-.97-1.042-2.869 2.54-1.61S19.9 16.33 24 16.33"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.426 24.579c0 .02 4.863-1.514 4.863-1.514m-4.863 4.488H43.5m-5.074 3.157l4.845 1.423"/>`),
		g.Group(children),
	)
}