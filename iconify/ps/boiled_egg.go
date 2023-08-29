package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoiledEgg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M3 256q0 49 29.5 89t74.5 54q-5 27-16 48.5T67 469H45q-9 0-15 6t-6 16q0 9 6 15t15 6h214q9 0 15-6t6-15q0-10-6-16t-15-6h-22q-13 0-24-21.5T197 399q48-14 76-58t28-106v-32q0-62-40-126Q214 0 152 0T43 77Q3 141 3 203v53zm175 213h-49q17-30 21-64h4q10 41 24 64zm-26-106q-45 0-76-31t-31-76h212q-4 45-30 76t-75 31zM45 203q0-50 34-105q36-55 73-55t73 55q34 55 34 105v10H45v-10z"/>`),
		g.Group(children),
	)
}