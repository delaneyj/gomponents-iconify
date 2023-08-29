package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paypal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1018 434q-25 121-122 195.5T672 704H480q-38 0-64 32l-61 254q-3 14-15.5 24t-26.5 10H155q-14 0-21.5-10t-4.5-24l7-30h113q14 0 26.5-10t15.5-24l61-254q26-32 64-32h192q127 0 224-74.5T954 370q5-23 6-49V213l2.5 2l3.5 3q36 42 50.5 98t1.5 118zM544 576H352q-38 0-64 32l-61 254q-3 14-15.5 24T185 896H27q-14 0-21.5-10T1 862L191 42q3-13 22.5-27.5T248 0q302 3 339 3q64 0 142 25.5T838 90q36 42 50.5 98t1.5 118q-25 121-122 195.5T544 576zm0-384h-96q-32 0-42.5 6.5T384 224l-32 122q-4 13 12.5 25.5T395 384q10 0 43-1t41-1q72 0 115-30.5t43-78.5q0-41-21.5-61T544 192z"/>`),
		g.Group(children),
	)
}