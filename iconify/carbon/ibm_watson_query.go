package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmWatsonQuery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 24h-.5l-3.3-6c.5-.5.8-1.2.8-2c0-1.7-1.3-3-3-3s-3 1.3-3 3c0 .8.3 1.5.8 2l-3.3 6H10c-1.7 0-3 1.3-3 3s1.3 3 3 3s3-1.3 3-3c0-.8-.3-1.5-.8-2l3.3-6.1h1l3.3 6.1c-.5.5-.8 1.2-.8 2c0 1.7 1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3zm-6-9c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1zm-6 13c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1zm12 0c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1z"/><path fill="currentColor" d="M24.8 9.1C23.9 5 20.3 2 16 2S8.1 5 7.2 9.1c-3 .6-5.2 3.3-5.2 6.4C2 19.1 4.9 22 8.5 22H9v-2h-.5C6 20 4 18 4 15.5c0-2.3 1.8-4.3 4.1-4.5H9l.1-.8C9.5 6.6 12.5 4 16 4c3.5 0 6.5 2.6 6.9 6.1l.1.9l.8.1c2.3.2 4.1 2.2 4.1 4.5c0 2.5-2 4.5-4.5 4.5H23v2h.5c3.6 0 6.5-2.9 6.5-6.5c0-3.2-2.2-5.9-5.2-6.5z"/>`),
		g.Group(children),
	)
}