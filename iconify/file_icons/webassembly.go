package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webassembly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m394.252 320.607l16.984 66.557h-46.052l14.796-66.557h14.272zM314.612 0H512v512H0V0h197.388c0 81.883 117.224 81.883 117.224 0zM152.667 275.934h-33.92l39.059 181.26h34.348l27.094-123.361l25.191 123.361h33.683l43.198-181.26h-33.255l-26.69 124.931l-25.356-124.931h-31.733l-28.188 123.409l-23.43-123.41zm313.756 181.26l-53.331-181.26H359.57l-43.864 181.26h34.111l8.897-40.344h60.943l11.56 40.344h35.206z"/>`),
		g.Group(children),
	)
}