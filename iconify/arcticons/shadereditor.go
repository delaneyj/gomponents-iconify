package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shadereditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2z"/><path fill="none" stroke="currentColor" d="m5.5 17.2l37-8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.7 18.9l-2 5.3l-2-5.3m9.5 4.3c-.3.6-1 1-1.7 1h0c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2h0c1.1 0 2 .9 2 2v.7h-4m9.4 1.6c-.3.6-1 1-1.7 1h0c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2h0c.7 0 1.4.4 1.7 1m11.4-4.8c-1.2 1.4-1.9 3.2-2 5.1c.1 1.9.8 3.7 2 5.1M19.2 37.9h4m-4-6.9l2-1.1m0 0v8"/><circle cx="26" cy="37.9" r=".8" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.8 27.8v12.5m2.5-20.1c1.1 0 2-.9 2-2h0c0-1.1-.9-2-2-2m0 8c1.1 0 2-.9 2-2h0c0-1.1-.9-2-2-2M28 23.5c.7.5 1.6.8 2.5.7h.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28 16.9c.7-.5 1.6-.8 2.5-.7h.8m-2 4h2"/>`),
		g.Group(children),
	)
}