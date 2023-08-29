package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeClimate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m342.766 113.36l-97.78 97.242l58.023 58.023l39.757-39.756l111.748 112.285L512 283.131L342.766 113.36zm-172.995 57.486L0 340.617l58.023 58.023l111.748-111.748L281.52 398.64l58.023-58.023l-169.77-169.771z"/>`),
		g.Group(children),
	)
}