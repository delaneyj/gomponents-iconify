package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerLowVolume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8.012 9.013V9h-4.21C2.808 9 2 9.847 2 10.892v10.216C2 22.153 2.807 23 3.803 23h4.195l8.613 3.83a1.693 1.693 0 0 0 2.425-1.525v-6.228a3.055 3.055 0 0 0 0-6.107V6.497c0-1.242-1.866-1.794-2.989-1.265l-8.035 3.78Zm9.024 15.818l-7.038-3.13V10.29l6.893-3.244h.002a.384.384 0 0 1 .114-.013h.029v17.8Z"/>`),
		g.Group(children),
	)
}