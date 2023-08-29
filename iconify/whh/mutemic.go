package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mutemic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1007.81 1007.5q-16.5 16.5-40 16.5t-39.5-17l-911-911q-17-16-17-39.5t16.5-40t40-16.5t39.5 16l911 911q17 17 17 40.5t-16.5 40zM732.31 579q36-61 36-131q0-26 19-45t45.5-19t45 19t18.5 45q0 124-71 224zm-409-409q8-72 62-120.5t127-48.5q80 0 136 56t56 135v256q0 43-19 84zm162 467q-63-9-108-53.5t-54-107.5zm-292.5-253q26.5 0 45.5 19t19 45q0 106 74.5 181t180.5 75q17 0 37-3l103 104q-38 15-76 21v70h128q27 0 45.5 19t18.5 45.5t-18.5 45t-45.5 18.5h-383q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h127v-70q-137-23-228-129.5t-91-248.5q0-26 18.5-45t45-19z"/>`),
		g.Group(children),
	)
}