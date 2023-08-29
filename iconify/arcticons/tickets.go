package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tickets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.337 43.5H14.663a2 2 0 0 1-2-2v-35a2 2 0 0 1 2-2h18.674a2 2 0 0 1 2 2v35a2 2 0 0 1-2 2ZM21.098 6.042h3.628"/><circle cx="26.553" cy="6.087" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.304 14.924H15.5v3.804m17 0v-3.804h-3.804m0 17H32.5V28.12m-14.68-3.589v5.054h5.054v-5.054Zm0-7.287v5.054h5.054v-5.054Zm7.287 0v5.054h5.054v-5.054Zm2.536 12.36a2.66 2.66 0 0 0 1.595-.965a3.845 3.845 0 0 0 .942-2.628v-1.48h-5.054v1.48h0a3.854 3.854 0 0 0 .942 2.628a2.66 2.66 0 0 0 1.575.965ZM15.5 28.12v3.804h3.804M31.051 4.5h-.603a1 1 0 0 0-.864.496l-.982 1.686a2 2 0 0 1-1.728.992h-5.748a2 2 0 0 1-1.728-.992l-.982-1.686a1 1 0 0 0-.864-.496h-.603"/>`),
		g.Group(children),
	)
}