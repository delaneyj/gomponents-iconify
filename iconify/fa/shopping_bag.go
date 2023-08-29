package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m1757 1408l35 313q3 28-16 50q-19 21-48 21H64q-29 0-48-21q-19-22-16-50l35-313h1722zm-93-839l86 775H42l86-775q3-24 21-40.5t43-16.5h256v128q0 53 37.5 90.5T576 768t90.5-37.5T704 640V512h384v128q0 53 37.5 90.5T1216 768t90.5-37.5T1344 640V512h256q25 0 43 16.5t21 40.5zm-384-185v256q0 26-19 45t-45 19t-45-19t-19-45V384q0-106-75-181t-181-75t-181 75t-75 181v256q0 26-19 45t-45 19t-45-19t-19-45V384q0-159 112.5-271.5T896 0t271.5 112.5T1280 384z"/>`),
		g.Group(children),
	)
}