package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeepassTwoAndroidoffline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.33 30.79a7.37 7.37 0 1 0 7.37 7.36h0a7.36 7.36 0 0 0-7.35-7.36Zm3 10.24l-5.42.08a2.67 2.67 0 0 1-.11-5.33h.2l-1.13-1.16l7.63 7.53ZM36 35.8a2.84 2.84 0 0 1 3.86-1A2.89 2.89 0 0 1 41 35.88a4.78 4.78 0 0 1 .38.87a2.57 2.57 0 0 1 2.2 2.44A1.66 1.66 0 0 1 42.12 41h-.77m-28.2-20a3.07 3.07 0 1 1 3.06 3.07A3.07 3.07 0 0 1 13.15 21Zm18.64 3.06A3.07 3.07 0 1 1 34.86 21a3.08 3.08 0 0 1-3.07 3.07ZM24 30.79a2.61 2.61 0 0 1 1.59 4.69l1.24 4.31h-5.67l1.23-4.31A2.62 2.62 0 0 1 24 30.79Zm18.5 1.28v-3.44"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.51 28.63v11.73a1.94 1.94 0 0 0 1.94 2h24.7M39 14l2.58-2.56a3 3 0 1 0-4.28-4.26h0l-2.55 2.6a18.39 18.39 0 0 0-21.51 0l-2.56-2.53a3.22 3.22 0 0 0-2.18-.92a3 3 0 0 0-2.12 5.15L9 14a18.41 18.41 0 0 0-3.47 10.8v3.8h37v-3.8A18.57 18.57 0 0 0 39 14Zm-19.86 7a2.93 2.93 0 1 1-2.93-2.92A2.93 2.93 0 0 1 19.14 21Zm12.65 2.92A2.92 2.92 0 1 1 34.71 21a2.92 2.92 0 0 1-2.92 2.94ZM5.51 35.29h13.77m9.42 0h2.85"/>`),
		g.Group(children),
	)
}