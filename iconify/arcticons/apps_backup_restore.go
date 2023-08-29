package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppsBackupRestore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.506 28.31v11.732a1.947 1.947 0 0 0 1.947 1.947h33.1a1.947 1.947 0 0 0 1.947-1.947V28.31m-3.495-14.593l2.59-2.55A3.022 3.022 0 1 0 37.313 6.9l-.002.001l-2.551 2.6a18.41 18.41 0 0 0-21.515 0l-2.57-2.57a3.174 3.174 0 0 0-2.17-.916a2.978 2.978 0 0 0-2.123.877a3.018 3.018 0 0 0 0 4.264l2.6 2.55a18.497 18.497 0 0 0-3.476 10.807v3.796H42.5v-3.796a18.497 18.497 0 0 0-3.495-10.797Zm-19.87 6.98a2.92 2.92 0 1 1-2.92-2.92h0a2.92 2.92 0 0 1 2.92 2.92v0Zm12.656 2.92a2.92 2.92 0 1 1 2.92-2.92a2.92 2.92 0 0 1-2.92 2.92h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.99 33.864h-2.191v-2.179"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.351 37.974a4.031 4.031 0 0 0 5.68.002a3.98 3.98 0 0 0 .002-5.65a4.031 4.031 0 0 0-5.68-.003l-.002.002l-1.552 1.54"/>`),
		g.Group(children),
	)
}