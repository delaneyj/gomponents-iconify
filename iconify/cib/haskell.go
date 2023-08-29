package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Haskell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M.011 27.292L7.531 16L.011 4.708h5.645L13.219 16L5.656 27.292zm7.52 0L15.052 16L7.531 4.708h5.688l15.005 22.584h-5.631l-4.713-7.12l-4.661 7.12zm18.204-6.595l-2.505-3.765h8.76v3.765zm-3.771-5.645l-2.5-3.765h12.515v3.765z"/>`),
		g.Group(children),
	)
}