package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tabs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M430 96H108c-19 0-36 17-36 36v286c0 19 17 36 36 36h35v72h-35C49 526 0 477 0 418V132C0 73 49 24 108 24h322c59 0 108 49 108 108h-72c0-19-17-36-36-36zm-143 72h322c59 0 108 48 108 107v287c0 59-49 107-108 107H287c-59 0-108-48-108-107V275c0-59 49-107 108-107zm0 430h322c19 0 36-17 36-36V275c0-19-17-36-36-36H287c-19 0-36 17-36 36v287c0 19 17 36 36 36z"/>`),
		g.Group(children),
	)
}