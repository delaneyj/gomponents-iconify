package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialWindows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M85.905 47.797V13l-39.97 5.83v28.967zM43.121 19.241l-29.026 4.234v24.322h29.026zM14.095 51.774v24.632l29.026 4.283V51.774zm31.84 29.331L85.905 87V51.774h-39.97z"/>`),
		g.Group(children),
	)
}