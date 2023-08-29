package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M170.671 384.805h243.816l366.784-298.94v1028.27l-366.784-298.94H170.671v-430.39zm748.41-48.764c72.085 72.084 108.835 159.717 110.248 262.898c0 98.938-36.749 183.744-110.248 254.417l-74.205-76.325c50.884-50.884 76.325-110.954 76.325-180.212c0-70.672-25.441-132.156-76.325-184.453l74.205-76.325z"/>`),
		g.Group(children),
	)
}