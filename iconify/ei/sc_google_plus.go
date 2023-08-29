package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScGooglePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 23v4.8h7.9c-.3 2.1-2.4 6-7.9 6c-4.8 0-8.7-4-8.7-8.8s3.9-8.8 8.7-8.8c2.7 0 4.5 1.2 5.6 2.2l3.8-3.7C24.9 12.4 21.8 11 18 11c-7.7 0-14 6.3-14 14s6.3 14 14 14c8.1 0 13.4-5.7 13.4-13.7c0-.9-.1-1.6-.2-2.3H18zm30 0h-4v-4h-4v4h-4v4h4v4h4v-4h4z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}