package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Temperature(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M320 1024q-87 0-160.5-43T43 864.5T0 704q0-77 34.5-144T128 448V197q0-81 56-139T319.5 0t136 58T512 197v251q59 45 93.5 112T640 704q0 87-43 160.5T480.5 981T320 1024zm64-501V192q0-27-19-45.5T319.5 128t-45 18.5T256 192v331q-57 20-92.5 69.5T128 704q0 80 56 136t135.5 56t136-56T512 704q0-62-35.5-111.5T384 523z"/>`),
		g.Group(children),
	)
}