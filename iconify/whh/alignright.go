package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alignright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 704H128q-27 0-45.5-18.5t-18.5-45T82.5 595t45.5-19h832q26 0 45 19t19 45.5t-19 45t-45 18.5zm0-192H192q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h768q26 0 45 19t19 45.5t-19 45t-45 18.5zm0-192H64q-27 0-45.5-18.5T0 256.5T18.5 211T64 192h896q26 0 45 19t19 45.5t-19 45t-45 18.5zm0-192H256q-27 0-45.5-18.5t-18.5-45T210.5 19T256 0h704q26 0 45 19t19 45.5t-19 45t-45 18.5zM576 768h384q26 0 45 19t19 45.5t-19 45t-45 18.5H576q-27 0-45.5-18.5T512 832t18.5-45.5T576 768z"/>`),
		g.Group(children),
	)
}