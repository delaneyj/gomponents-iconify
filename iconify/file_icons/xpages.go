package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xpages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m297.164 247.455l187.596-187.6h-92.476l-142.478 142.48l-149.541-142.48H5.458l198.096 188.733L0 452.145h92.476l158.435-158.437l166.296 158.437H512z"/>`),
		g.Group(children),
	)
}