package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandVimeo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 8.5l1 1S5.5 8.398 6 9c.509.609 1.863 7.65 2.5 9c.556 1.184 1.978 2.89 4 1.5C14.5 18 20 14 21 8c.444-2.661-1-4-2.5-4c-2 0-4.047 1.202-4.5 4c2.05-1.254 2.551 1 1.5 3c-1.052 2-2 3-2.5 3c-.49 0-.924-1.165-1.5-3.5c-.59-2.42-.5-6.5-3-6.5S3 8.5 3 8.5z"/>`),
		g.Group(children),
	)
}