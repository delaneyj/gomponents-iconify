package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M462 192q0 43-30 72.5T359 294t-73-29.5t-30-72.5t30-72.5T359 90t73 29.5t30 72.5zM105 90q-43 0-73 29.5T2 192t30 72.5t73 29.5t73-29.5t30-72.5t-30-72.5T105 90z"/>`),
		g.Group(children),
	)
}