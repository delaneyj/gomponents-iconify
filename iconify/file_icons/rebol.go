package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rebol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M508.001 256c0 141.385-112.825 256-252.001 256S3.999 397.385 3.999 256S116.824 0 256 0s252.001 114.615 252.001 256zM248.975 117.092c-75.701 0-137.069 60.555-137.069 135.253s61.368 135.253 137.069 135.253s137.068-60.555 137.068-135.253s-61.368-135.253-137.068-135.253zm0 79.956c-30.95 0-56.04 24.757-56.04 55.297s25.09 55.297 56.04 55.297s56.039-24.757 56.039-55.297s-25.09-55.297-56.04-55.297z"/>`),
		g.Group(children),
	)
}