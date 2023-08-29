package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M990 1L856 135q-69-63-157.5-98.5T512 1Q353 1 223.5 90.5T37 323l119 48q43-108 139.5-175T512 129q145 0 254 97L640 353q-1 14 8.5 23.5T672 385h309q14 0 27.5-13.5T1023 344l1-320q1-24-34-23zM512 897q-145 0-254-96l126-127q1-14-8.5-23.5T352 641H43q-14 1-27.5 14.5T1 683l-1 320q-1 24 34 23l134-134q69 63 157.5 98t186.5 35q159 0 288.5-89T987 703l-119-47q-43 108-139.5 174.5T512 897z"/>`),
		g.Group(children),
	)
}