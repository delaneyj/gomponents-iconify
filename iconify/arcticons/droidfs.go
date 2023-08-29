package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Droidfs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.506 27.25v12.79a1.947 1.947 0 0 0 1.947 1.947h33.1A1.947 1.947 0 0 0 42.5 40.04V27.25Zm33.499-13.534l2.59-2.551a3.022 3.022 0 0 0-4.282-4.266l-.002.002L34.76 9.5a18.41 18.41 0 0 0-21.515 0l-2.57-2.57a3.174 3.174 0 0 0-2.17-.915a2.979 2.979 0 0 0-2.123.876a3.018 3.018 0 0 0 0 4.264l2.6 2.55a18.497 18.497 0 0 0-3.476 10.807v2.738H42.5v-2.738a18.497 18.497 0 0 0-3.495-10.796Zm-19.87 6.98a2.92 2.92 0 1 1-2.92-2.921a2.92 2.92 0 0 1 2.92 2.92Zm12.656 2.92a2.92 2.92 0 1 1 2.92-2.92a2.92 2.92 0 0 1-2.92 2.92Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.367 33.614v-1.4a2.633 2.633 0 0 1 5.266 0v1.4m-6.63 0h7.994v5.967h-7.994zM24 37.209v-1.223"/>`),
		g.Group(children),
	)
}