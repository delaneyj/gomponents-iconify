package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListRotateTwoHundredSeventyTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.5 20.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm.5-4V2.997a1 1 0 0 1 1.993-.117l.007.117V16.5a1 1 0 0 1-1.993.116L17 16.5Zm-6.5 4a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm.5-4V2.997a1 1 0 0 1 1.993-.117l.007.117V16.5a1 1 0 0 1-1.993.116L11 16.5Zm-6.492 4a1.5 1.5 0 1 1 2.999 0a1.5 1.5 0 0 1-3 0Zm.493-4V2.997a1 1 0 0 1 1.993-.117l.007.117V16.5a1 1 0 0 1-1.993.116L5 16.5Z"/>`),
		g.Group(children),
	)
}