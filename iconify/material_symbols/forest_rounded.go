package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForestRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 22q-.825 0-1.413-.588T13 20v-1h4v1q0 .825-.588 1.413T15 22Zm-6 0q-.825 0-1.413-.588T7 20v-2H1.825q-.6 0-.875-.525T1 16.45L3.85 12q-.6 0-.85-.537t.1-1.038l5.075-7.25q.3-.425.825-.425t.825.425l5.075 7.25q.35.5.1 1.038t-.85.537L17 16.45q.275.425.013.988t-.838.562H11v2q0 .825-.588 1.413T9 22Zm8.875-4q.25-.25.338-.838T18 16.1l-2.375-3.725q.375-.35.475-1.087t-.3-1.313l-3.175-4.55l1.55-2.25q.3-.425.825-.425t.825.425l5.075 7.25q.35.5.1 1.038t-.85.537L23 16.45q.325.5.05 1.025t-.875.525h-4.3Z"/>`),
		g.Group(children),
	)
}