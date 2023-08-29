package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LastpassAuthenticator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5a.75.75 0 0 1 .26 0l16.08 4.82a.93.93 0 0 1 .66.88V26a15 15 0 0 1-4.87 11.2c-3 2.84-7.17 4.83-11.87 6.22a.86.86 0 0 1-.52 0c-4.7-1.39-8.85-3.38-11.87-6.22A15 15 0 0 1 7 26V10.2a.93.93 0 0 1 .66-.88l16.08-4.78a.75.75 0 0 1 .26 0Zm-8.91 15.2a2.87 2.87 0 1 1-2.87 2.87a2.87 2.87 0 0 1 2.87-2.87Zm7.3 0a2.87 2.87 0 1 1-2.87 2.87a2.88 2.88 0 0 1 2.87-2.87Zm7.29 0a2.87 2.87 0 1 1-2.87 2.87a2.87 2.87 0 0 1 2.87-2.87Zm5.44-4.6v15"/>`),
		g.Group(children),
	)
}