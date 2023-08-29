package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Librivox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.764 25.966c0-4.489-1.936-5.99-4.764-5.99s-4.764 1.501-4.764 5.99m-.55 0h.55v5.324h0h-.55a1.826 1.826 0 0 1-1.826-1.826v-1.672a1.826 1.826 0 0 1 1.826-1.826Zm10.628 5.324h-.55h0v-5.324h.55a1.826 1.826 0 0 1 1.826 1.826v1.672a1.826 1.826 0 0 1-1.826 1.826Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 42.476L12.938 33.58V5.524L24 16.351"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 42.476L9.125 35.557V7.501L24 16.351"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 42.476L5.5 38.004V9.948L24 16.351m0 26.125l11.062-8.896V5.524L24 16.351"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 42.476l14.875-6.919V7.501L24 16.351"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 42.476l18.5-4.472V9.948L24 16.351"/>`),
		g.Group(children),
	)
}