package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundAmplifier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="20.598" cy="19.108" r="4.785" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.157 19.108a12.44 12.44 0 1 1 24.88 0c0 3.73-1.394 7.393-4.24 9.356C21.121 33.757 21.66 42.5 14.642 42.5c-6.495 0-6.486-5.848-6.486-5.848m26.571-4.477A19.245 19.245 0 0 0 34.206 5.5"/>`),
		g.Group(children),
	)
}