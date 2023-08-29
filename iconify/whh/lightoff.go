package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightoff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1010.01 1010.5q-14.5 14.5-34.5 14.5t-35-14l-926-927q-14-14-14-34.5t14.5-35T49.51 0t35 15l926 926q14 14 14 34.5t-14.5 35zM668.51 515q72-82 108-144t20-78.5t-78 20t-144 108.5l-47-47q56-50 106.5-87.5t82.5-54.5t56-27t34-11l11-1q53 0 98.5 26t45.5 70v320q0 64-54 145zm-247 60q-72 82-108 143.5t-20 78t78-19.5t143-109l240 240q-81 53-145 53h-321q-44 0-70-45.5t-26-98.5q0-4 1.5-11t11-33t26.5-56.5t54.5-81.5t88.5-107z"/>`),
		g.Group(children),
	)
}