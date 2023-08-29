package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Try(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1152 704q0 191-94.5 353T801 1313.5T448 1408H288q-14 0-23-9t-9-23V765L41 831q-3 1-9 1q-10 0-19-6q-13-10-13-26V672q0-23 23-31l233-71v-93L41 543q-3 1-9 1q-10 0-19-6q-13-10-13-26V384q0-23 23-31l233-71V32q0-14 9-23t23-9h160q14 0 23 9t9 23v181L855 97q15-5 28 5t13 26v128q0 23-23 31L480 408v93l375-116q15-5 28 5t13 26v128q0 23-23 31L480 696v487q188-13 318-151t130-328q0-14 9-23t23-9h160q14 0 23 9t9 23z"/>`),
		g.Group(children),
	)
}