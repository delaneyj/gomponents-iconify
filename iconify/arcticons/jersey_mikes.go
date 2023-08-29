package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JerseyMikes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.256 19.442v2.7c0 1.1-.9 2-2 2c-.6 0-1.1-.2-1.4-.6m-25.112-10.2v6c0 1.1-.9 2-2 2s-2-.9-2-2v-.7m9.98 1.7c-.3.6-1 1-1.7 1c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2s2 .9 2 2v.7h-4m19.651 1.6c-.3.6-1 1-1.7 1c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2s2 .9 2 2v.7h-4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.256 16.042v3.3c0 1.1-.9 2-2 2s-2-.9-2-2v-3.3m-15.951-.1v5.4m0-3.4c0-1.1.9-2 2-2m2.09 5c.4.3.8.4 1.6.4h.4c.7 0 1.3-.6 1.3-1.3s-.6-1.3-1.3-1.3h-.9c-.7 0-1.3-.6-1.3-1.3s.6-1.3 1.3-1.3h.4c.9 0 1.3.1 1.6.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.501v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.303 34.259c.4.3.8.4 1.6.4h.4c.7 0 1.3-.6 1.3-1.3s-.6-1.3-1.3-1.3h-.9c-.7 0-1.3-.6-1.3-1.3s.6-1.3 1.3-1.3h.4c.9 0 1.3.1 1.6.4m-19.705-.6v5.4m12.658-1c-.3.6-1 1-1.7 1c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2s2 .9 2 2v.7h-4m-20.259 2.5v-7.9l4 8l4-8v8m7.641-2.9l2.8 2.8m-4-1.6l3.6-3.7m-3.6-2.7v8.1"/><circle cx="19.698" cy="26.859" r=".7" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.768 28.609c.5-.2.5-.7.5-1v-.7"/><circle cx="35.018" cy="26.909" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}