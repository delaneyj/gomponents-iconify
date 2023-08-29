package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ban(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1312 797q0-161-87-295l-754 753q137 89 297 89q111 0 211.5-43.5T1153 1184t116-174.5t43-212.5zm-999 299l755-754q-135-91-300-91q-148 0-273 73T297 523t-73 274q0 162 89 299zm1223-299q0 157-61 300t-163.5 246t-245 164t-298.5 61t-298.5-61t-245-164T61 1097T0 797t61-299.5T224.5 252t245-164T768 27t298.5 61t245 164T1475 497.5t61 299.5z"/>`),
		g.Group(children),
	)
}