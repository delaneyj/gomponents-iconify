package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webtoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.361 38.149l13.418-1.658v5.21l15.708-7.282L42.5 21.336h-4.795l1.184-15.037l-30.311 4.677l2.605 9.768H5.5l5.861 17.405zm.159-13.082h4.924M13.982 32.5v-7.433"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.556 32.5v-7.433L36.48 32.5v-7.433m-18.876 4.971a2.462 2.462 0 0 0 4.924 0V27.53a2.462 2.462 0 1 0-4.924 0Zm6.976 0a2.462 2.462 0 0 0 4.924 0V27.53a2.462 2.462 0 1 0-4.924 0Zm-1.159-7.105h3.716M23.421 15.5h3.716m-3.716 3.716h2.423M23.421 15.5v7.433M21.446 15.5l-1.858 7.433l-1.859-7.433l-1.858 7.433l-1.858-7.433"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M32.129 19.216a1.858 1.858 0 0 1 0 3.717h-3.066V15.5h3.066a1.858 1.858 0 0 1 0 3.716Zm0 0h-3.066"/>`),
		g.Group(children),
	)
}