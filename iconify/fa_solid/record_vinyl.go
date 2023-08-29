package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordVinyl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 152a104 104 0 1 0 104 104a104 104 0 0 0-104-104zm0 128a24 24 0 1 1 24-24a24 24 0 0 1-24 24zm0-272C119 8 8 119 8 256s111 248 248 248s248-111 248-248S393 8 256 8zm0 376a128 128 0 1 1 128-128a128 128 0 0 1-128 128z"/>`),
		g.Group(children),
	)
}