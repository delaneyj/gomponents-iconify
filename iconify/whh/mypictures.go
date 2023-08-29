package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mypictures(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.38 896h-768q-53 0-90.5-37.5T.38 768V256q0-26 19-45t45-19h480l46-84q7-19 28-31.5t44-12.5h240q24 0 45 12.5t29 31.5l48 84v576q0 53-37.5 90.5t-90.5 37.5zm-64-576h-640v394l91-92q15-15 37.5-15t37.5 15l145 146h64l-86-87l154-155q15-15 37.5-15t36.5 15l123 123V320zm-416 256q-40 0-68-28t-28-68t28-68t68-28t68 28t28 68t-28 68t-68 28zm-370-532q7-19 28-31.5t44-12.5h240q24 0 45 12.5t29 31.5l48 84H.38z"/>`),
		g.Group(children),
	)
}