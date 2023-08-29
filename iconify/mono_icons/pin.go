package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 10.6V4a1 1 0 0 1 0-2h12a1 1 0 1 1 0 2v6.6c.932 1.02 1.432 2.034 1.699 2.834c.146.438.22.81.26 1.08a4.43 4.43 0 0 1 .04.43v.034l.001.013v.008s-.005-.131 0 .001a1 1 0 0 1-1 1h-6v5a1 1 0 1 1-2 0v-5H5a1 1 0 0 1-1-1v-.022a2.013 2.013 0 0 1 .006-.134a5.07 5.07 0 0 1 .035-.33c.04-.27.114-.642.26-1.08c.267-.8.767-1.814 1.699-2.835zM16 4H8v7a1 1 0 0 1-.293.707c-.847.847-1.271 1.678-1.486 2.293H17.78c-.215-.615-.64-1.446-1.486-2.293A1 1 0 0 1 16 11V4z"/>`),
		g.Group(children),
	)
}