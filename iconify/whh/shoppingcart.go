package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shoppingcart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M961 128h-76L763 723q-5 19-21.5 32T707 768H160q-19 0-35.5-13T103 723L5 346q-10-36 3-63t40-27l709-123l18-89q5-18 21.5-31T831 0h130q27 0 45.5 18.5t18.5 45t-19 45.5t-45 19zM289 832q40 0 68 28t28 68t-28 68t-68 28t-68-28t-28-68t28-68t68-28zm320 0q40 0 68 28t28 68t-28 68t-68 28t-68-28t-28-68t28-68t68-28z"/>`),
		g.Group(children),
	)
}