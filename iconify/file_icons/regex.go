package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Regex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M309.677 349.209V230.914l-97.438 59.71l-34.351-59.84L279.57 175.14l-101.682-56.47l34.525-58.75l97.264 59.447V.618h69.421v118.75l98.161-59.446L512 118.67l-102.14 56.47L512 230.785l-34.878 60.112l-98.024-59.982V349.21h-69.42zm-164.504 89.587c0-55.67-60.68-90.652-108.962-62.817c-48.281 27.835-48.281 97.8 0 125.635c48.282 27.835 108.962-7.148 108.962-62.818z"/>`),
		g.Group(children),
	)
}