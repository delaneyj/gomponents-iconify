package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M64 72.5v316.3c0 2.8 1.3 5.5 3.5 7l69.8 50.2c5.2 3.7 12.1-.3 12.1-7V122.7c0-2.8-1.3-5.5-3.5-7L76.1 65.4c-5.2-3.7-12.1.3-12.1 7.1z" fill="currentColor"/><path d="M168.7 123.3v316.3c0 6.7 6.9 10.8 12.1 7l62-36.7c2.2-1.6 3.5-4.2 3.5-7V86.6c0-6.7-6.9-10.8-12.1-7l-62 36.7c-2.2 1.5-3.5 4.1-3.5 7z" fill="currentColor"/><path d="M435.9 65.5l-69.8 50.7c-2.2 1.6-3.5 4.2-3.5 7v316.2c0 6.8 6.9 10.8 12.1 7l69.8-50.7c2.2-1.6 3.5-4.2 3.5-7V72.5c0-6.7-6.9-10.7-12.1-7z" fill="currentColor"/><path d="M265.7 85.6v316.2c0 2.8 1.3 5.5 3.5 7l62 37c5.2 3.8 12.1-.3 12.1-7V122.6c0-2.8-1.3-5.5-3.5-7l-62-37c-5.2-3.7-12.1.3-12.1 7z" fill="currentColor"/>`),
		g.Group(children),
	)
}