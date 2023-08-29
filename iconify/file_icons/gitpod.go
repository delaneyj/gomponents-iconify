package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitpod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M221.995 0L444 128.07v255.973L221.995 512L0 384.019V128.03L221.995 0zm177.83 358.735V205.197L265.812 281.47v51.386l89.35-50.84v50.815l-133.168 76.223l-133.018-76.198V179.318l133.043-76.77l133.018 76.67l44.786-25.792L221.995 50.99L44.315 153.513v205.222l177.68 101.78l177.83-101.78z"/>`),
		g.Group(children),
	)
}