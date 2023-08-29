package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Help(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0zm-27 470h24c3-61 17-81 74-115c35-22 46-29 59-43c18-19 27-43 27-71c0-71-57-114-150-114c-32 0-57 6-82 15c-50 19-81 59-81 103c0 22 15 33 46 33c26 0 40-11 39-29v-24c0-15 6-41 14-54c10-16 31-27 55-27c45 0 74 36 74 89c0 20-3 39-11 57c-7 14-15 25-39 51c-38 40-49 69-49 114v15zm13 145c29 0 52-23 52-52s-22-52-52-52c-29 0-51 23-51 52s22 52 51 52z"/>`),
		g.Group(children),
	)
}