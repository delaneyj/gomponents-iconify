package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialFlickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M79 14H21c-3.85 0-7 3.15-7 7v58c0 3.85 3.15 7 7 7h58c3.85 0 7-3.15 7-7V21c0-3.85-3.15-7-7-7zM35.435 61.5c-6.456 0-11.69-5.233-11.69-11.689s5.234-11.69 11.69-11.69c6.456 0 11.689 5.233 11.689 11.69c0 6.456-5.233 11.689-11.689 11.689zm29.633 0c-6.456 0-11.69-5.233-11.69-11.689s5.233-11.69 11.69-11.69c6.456 0 11.689 5.233 11.689 11.69c.001 6.456-5.233 11.689-11.689 11.689z"/>`),
		g.Group(children),
	)
}