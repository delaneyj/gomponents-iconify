package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BemTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#010101" d="M146.335 187.82h108.833v31.347H146.335V187.82Z"/><path fill="#F7B334" d="M130.075 187.82v-56.947H78.937v56.947h51.138"/><path fill="#010101" d="M108.047 62.171h-41.96L34.092 0H.262v187.82H79.2v-56.947h66.874V100.05c0-21.42-15.473-37.878-38.026-37.878Zm12.587 50.94h-11.538l-13.113-2.874v-2.874h24.651v5.747Z"/>`),
		g.Group(children),
	)
}