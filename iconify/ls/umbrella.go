package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M288 589V347c10-3 20-4 26-4s16 1 26 4v242c0 42-28 105-105 105c-76 0-104-63-104-105c0-14 11-26 26-26c14 0 26 12 26 26c0 13 5 53 52 53c50 0 53-44 53-53zm52-549v28c161 13 288 148 288 312h-4c-12-45-52-69-101-69s-89 24-101 69h-7c-10-36-38-55-75-65c-8-2-17-3-26-3s-18 1-26 3c-37 10-66 29-75 65h-8c-11-45-52-69-100-69c-49 0-90 24-101 69H0C0 216 127 81 288 68V40c0-14 11-26 26-26c14 0 26 12 26 26z"/>`),
		g.Group(children),
	)
}