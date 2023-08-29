package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaskSadLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M21 12V6.719c0-2.19 0-3.285-.707-3.884c-.707-.6-1.788-.42-3.95-.059l-1.055.176c-1.64.273-2.46.41-3.288.41c-.828 0-1.648-.137-3.288-.41l-1.054-.176c-2.162-.36-3.244-.54-3.95.059C3 3.434 3 4.529 3 6.719V12c0 5.49 4.239 8.155 6.899 9.286c.721.307 1.082.46 2.101.46c1.02 0 1.38-.153 2.101-.46C16.761 20.155 21 17.49 21 12Z"/><path stroke-linecap="round" d="M6.5 9c.291-.583 1.077-1 2-1s1.709.417 2 1m3 0c.291-.583 1.077-1 2-1s1.709.417 2 1m-9 6s1.05-1 3.5-1s3.5 1 3.5 1"/></g>`),
		g.Group(children),
	)
}