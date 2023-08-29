package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Walli(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="38.903" cy="17.514" r="4.597" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.616 32.703c.8-1.3.8-2.998 0-4.397l-7.594-13.09c-1.3-2.199-4.098-2.898-6.296-1.7c-.8.5-1.399 1.2-1.799 2c-.6 1.298-.6 2.797.2 4.196l7.495 13.09c1.3 2.2 4.097 2.899 6.295 1.7c.8-.5 1.4-1.1 1.7-1.799Zm14.89 0c.8-1.3.8-2.998 0-4.397l-7.495-13.09c-1.3-2.199-4.097-2.898-6.296-1.7c-.799.5-1.399 1.2-1.798 2c-.6 1.298-.6 2.797.2 4.196l7.494 13.09c1.3 2.2 4.097 2.899 6.296 1.7c.7-.5 1.299-1.1 1.599-1.799Z"/>`),
		g.Group(children),
	)
}