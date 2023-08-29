package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Goibibo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.207 21.752v5.562c0 1.023-.83 1.853-1.853 1.853h0c-.512 0-.976-.207-1.311-.543"/><rect width="3.707" height="4.912" x="8.5" y="21.752" rx="1.854" ry="1.854" transform="rotate(180 10.354 24.209)"/></g><rect width="3.707" height="4.912" x="14.39" y="21.752" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.854" ry="1.854"/><rect width="3.707" height="4.912" x="35.793" y="21.752" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.854" ry="1.854"/><circle cx="20.225" cy="19.482" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.225 21.752v4.913"/><circle cx="27.947" cy="19.482" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.947 21.752v4.913m-5.653-3.059c0-1.024.83-1.854 1.854-1.854h0c1.024 0 1.854.83 1.854 1.854v1.205c0 1.024-.83 1.854-1.854 1.854h0a1.854 1.854 0 0 1-1.854-1.854m0 1.854V19.25m7.773 4.356c0-1.024.83-1.854 1.854-1.854h0c1.024 0 1.854.83 1.854 1.854v1.205c0 1.024-.83 1.854-1.854 1.854h0a1.854 1.854 0 0 1-1.854-1.854m0 1.854V19.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}