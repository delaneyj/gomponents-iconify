package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FamilysearchTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.55 18.7h5.5v5.5h-5.5zm14.6-3.3h5.1v5.1h-5.1zm-6.7 1h4.3v4.3h-4.3zm2.9 6.7h5.4v5.4h-5.4zm7.8-.1h2.3v2.3h-2.3zm-15.8-9.5h2.9v2.9h-2.9zm5.1-3.6h4.3v4.3h-4.3z"/><path d="M22.95 20.8c-3.7 12 3.8 17.3 3.8 17.3h-3.4s-6.7-6.6-.4-17.3"/><path d="M25.55 28.5s-.7 2.3-2.9 3.3m-5.3-7.6s.4 1.9 2.7 3.5"/></g>`),
		g.Group(children),
	)
}