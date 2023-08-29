package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icebox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.535 33.562l-5.534-1.833l-5.535 1.833l-1.176-5.697L12.931 24l4.359-3.865l1.176-5.697L24 16.272l5.535-1.833l1.175 5.697l4.36 3.865l-4.359 3.864ZM24 31.73V2.5m-6.339 4.651l6.34 3.028m6.339-3.028L24 10.179m0 21.551V45.5m6.34-4.65L24 37.822m-6.339 3.028l6.34-3.028m-6.711-9.957L42.668 13.25M35.46 10.1l.54 6.99m5.8 3.962l-5.799-3.962M17.29 27.865L5.333 34.75m7.207 3.151l-.54-6.99M6.202 26.95L12 30.911m18.71-3.046l11.958 6.886M41.8 26.95l-5.798 3.962m-.542 6.989l.54-6.99m-5.29-3.046L5.334 13.25m.868 7.802L12 17.09m.54-6.99L12 17.09"/>`),
		g.Group(children),
	)
}