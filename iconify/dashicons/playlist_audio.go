package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistAudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 3V1H2v2h15zm0 4V5H2v2h15zm-7 4V9H2v2h8zm7.45-1.96l-6 1.12c-.16.02-.19.03-.29.13c-.11.09-.16.22-.16.37v4.59c-.29-.13-.66-.14-.93-.14c-.54 0-1 .19-1.38.57s-.56.84-.56 1.38c0 .53.18.99.56 1.37s.84.57 1.38.57c.49 0 .92-.16 1.29-.48s.59-.71.65-1.19v-4.95L17 11.27v3.48c-.29-.13-.56-.19-.83-.19c-.54 0-1.11.19-1.49.57c-.38.37-.57.83-.57 1.37s.19.99.57 1.37s.84.57 1.38.57c.53 0 .99-.19 1.37-.57s.57-.83.57-1.37V9.6c0-.16-.05-.3-.16-.41a.428.428 0 0 0-.39-.15zM8 15v-2H2v2h6zm-2 4v-2H2v2h4z"/>`),
		g.Group(children),
	)
}