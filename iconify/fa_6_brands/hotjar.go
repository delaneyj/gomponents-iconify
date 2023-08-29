package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hotjar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M361.5 0c0 131.6-80.7 176.8-140.2 209.4c-.6.3-1.1.6-1.6.9c-53.8 30.2-88.7 49.8-89.6 122H32c0-131.5 80.7-176.7 140.2-209.3C227 93.2 262.5 73 262.5 0h98.9zM301 302.6c54.8-29.8 90.3-50 90.3-123h98c0 131.6-80.7 176.7-140.2 209.4c-54.8 29.8-90.3 50-90.3 123h-98c0-131.6 80.7-176.8 140.2-209.4z"/>`),
		g.Group(children),
	)
}