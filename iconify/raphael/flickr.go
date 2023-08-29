package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.77 8.895a7.09 7.09 0 0 0-5.77 2.97a7.09 7.09 0 0 0-5.77-2.97A7.104 7.104 0 0 0 3.125 16a7.104 7.104 0 0 0 7.105 7.105c2.38 0 4.48-1.175 5.77-2.97a7.092 7.092 0 0 0 5.77 2.97a7.105 7.105 0 1 0 0-14.21zm0 12.927a5.826 5.826 0 0 1-5.822-5.82a5.827 5.827 0 0 1 5.82-5.825a5.828 5.828 0 0 1 5.825 5.824a5.828 5.828 0 0 1-5.824 5.822z"/>`),
		g.Group(children),
	)
}