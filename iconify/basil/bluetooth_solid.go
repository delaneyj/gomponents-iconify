package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.02 6.48a.75.75 0 0 1 1.06-1.06l4.67 4.669V2.604c0-.891 1.077-1.338 1.707-.708l4.346 4.347a1 1 0 0 1 0 1.414L13.51 11.95l4.293 4.293a1 1 0 0 1 0 1.414l-4.346 4.346c-.63.63-1.707.184-1.707-.707v-7.485l-4.67 4.67a.75.75 0 1 1-1.06-1.062l5.47-5.47L6.02 6.48Zm7.23-2.67v6.279l3.14-3.14l-3.14-3.138Zm3.14 13.14l-3.14-3.14v6.279l3.14-3.14Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}