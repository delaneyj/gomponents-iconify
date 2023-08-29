package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zM410.376 261.328c181.004.811 314.621 181.414 529.028 42.7v362.988c-173.222 173.222-335.647-107.109-588.721-38.232v309.888h-90.088V304.028c54.425-31.422 103.643-42.906 149.781-42.7z"/>`),
		g.Group(children),
	)
}