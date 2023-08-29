package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tunein(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 704H704V512q0-53 37.5-90.5T832 384h192v192q0 53-37.5 90.5T896 704zm-384 320H320V768h320v128q0 53-37.5 90.5T512 1024zM0 576V384h512q53 0 90.5 37.5T640 512v192H128q-53 0-90.5-37.5T0 576zm480-256q-66 0-113-47t-47-113t47-113T480 0t113 47t47 113t-47 113t-113 47z"/>`),
		g.Group(children),
	)
}