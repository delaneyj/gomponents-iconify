package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloudhosting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1024H128q-53 0-90.5-37.5T0 896t37.5-90.5T128 768h768q53 0 90.5 37.5T1024 896t-37.5 90.5T896 1024zM191.5 832q-26.5 0-45 18.5t-18.5 45t18.5 45.5t45 19t45.5-18.5t19-45.5t-19-45.5t-45.5-18.5zm640.5 0H448q-27 0-45.5 18.5t-18.5 45t18.5 45.5t45.5 19h384q26 0 45-18.5t19-45.5t-19-45.5t-45-18.5zm-96-128H256q-80 0-136-56T64 512q0-57 30.5-103.5T175 338q34 54 87.5 90T380 475q-57-34-90.5-92T256 256q0-106 75-181T512 0q81 0 146 45.5T751 164q-103 18-171 98t-68 186q0 11 1 22q4-89 68.5-151.5T736 256q93 0 158.5 65.5T960 480t-65.5 158.5T736 704z"/>`),
		g.Group(children),
	)
}