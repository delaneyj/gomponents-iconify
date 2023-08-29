package fa_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wpexplorer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M512 256c0 141.2-114.7 256-256 256C114.8 512 0 397.3 0 256S114.7 0 256 0s256 114.7 256 256zm-32 0c0-123.2-100.3-224-224-224C132.5 32 32 132.5 32 256s100.5 224 224 224s224-100.5 224-224zM160.9 124.6l86.9 37.1l-37.1 86.9l-86.9-37.1l37.1-86.9zm110 169.1l46.6 94h-14.6l-50-100l-48.9 100h-14l51.1-106.9l-22.3-9.4l6-14l68.6 29.1l-6 14.3l-16.5-7.1zm-11.8-116.3l68.6 29.4l-29.4 68.3L230 246l29.1-68.6zm80.3 42.9l54.6 23.1l-23.4 54.3l-54.3-23.1l23.1-54.3z"/>`),
		g.Group(children),
	)
}