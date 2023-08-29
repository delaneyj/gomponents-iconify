package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JumpRopeTrainingPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.902 18.598v14.098m28.196-14.098v14.098"/><circle cx="24" cy="18.598" r="14.098" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="18.598" r="9.078" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="18.598" r="4.058" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.532 32.696h4.74v8.434a2.37 2.37 0 0 1-2.37 2.37h0a2.37 2.37 0 0 1-2.37-2.37v-8.434h0Zm28.197 0h4.739v8.434a2.37 2.37 0 0 1-2.37 2.37h0a2.37 2.37 0 0 1-2.37-2.37v-8.434h0Z"/>`),
		g.Group(children),
	)
}