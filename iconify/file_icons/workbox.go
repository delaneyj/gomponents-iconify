package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Workbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m148.102 368.333l-23.354-7.095l-1.182-282.605l-53.802 25.718l-2.956 240.333l-19.214-5.913l-1.479-223.187L0 137.755v236.49L211.954 476.23V35.769L150.467 65.33l-2.365 303.003zm90.457-324.287v197.469L512 246.245v-117.95L238.559 44.047zm61.192 97.552l167.612 32.222v18.919L299.75 172.046v-30.448zm-61.192 128.887v197.764L512 383.704V265.755l-273.441 4.73zm61.192 67.99l167.612-19.51v18.92L299.75 368.924v-30.448z"/>`),
		g.Group(children),
	)
}