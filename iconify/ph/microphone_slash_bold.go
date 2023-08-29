package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneSlashBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M56.88 31.93a12 12 0 1 0-17.76 16.14L76 88.64V120a52 52 0 0 0 72.11 48l11.26 12.39A67.34 67.34 0 0 1 128 188a68.07 68.07 0 0 1-68-68a12 12 0 0 0-24 0a92.14 92.14 0 0 0 80 91.22V232a12 12 0 0 0 24 0v-20.77a90.39 90.39 0 0 0 35.92-12.68l23.2 25.52a12 12 0 0 0 17.76-16.14ZM128 148a28 28 0 0 1-28-28v-5l29.9 32.89c-.63.11-1.26.11-1.9.11ZM90.67 45.27a12 12 0 0 1-.48-17A52 52 0 0 1 180 64v48.54a12 12 0 1 1-24 0V64a28 28 0 0 0-48.36-19.22a12 12 0 0 1-16.97.49ZM191 145.58a67.63 67.63 0 0 0 5-25.58a12 12 0 0 1 24 0a91.48 91.48 0 0 1-6.74 34.61a12 12 0 0 1-22.23-9Z"/>`),
		g.Group(children),
	)
}