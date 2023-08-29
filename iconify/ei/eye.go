package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M25 36C13.5 36 8.3 25.9 8.1 25.4c-.1-.3-.1-.6 0-.9C8.3 24.1 13.5 14 25 14s16.7 10.1 16.9 10.6c.1.3.1.6 0 .9c-.2.4-5.4 10.5-16.9 10.5zM10.1 25c1.1 1.9 5.9 9 14.9 9s13.7-7.1 14.9-9c-1.1-1.9-5.9-9-14.9-9s-13.7 7.1-14.9 9z"/><path fill="currentColor" d="M25 34c-5 0-9-4-9-9s4-9 9-9s9 4 9 9s-4 9-9 9zm0-16c-3.9 0-7 3.1-7 7s3.1 7 7 7s7-3.1 7-7s-3.1-7-7-7z"/><path fill="currentColor" d="M25 30c-2.8 0-5-2.2-5-5c0-.6.4-1 1-1s1 .4 1 1c0 1.7 1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3c-.6 0-1-.4-1-1s.4-1 1-1c2.8 0 5 2.2 5 5s-2.2 5-5 5z"/>`),
		g.Group(children),
	)
}