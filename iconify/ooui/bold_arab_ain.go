package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoldArabAin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.89 11.9c0 1.57 1.61 2.44 4.85 2.63l2.55-.04l.37.05c-.03.14-.29.4-.77.76l-.1.07A6.97 6.97 0 0 1 9.63 17a4.3 4.3 0 0 1-3.16-1.16a4.3 4.3 0 0 1-1.13-3.15c0-1.58.66-3 1.96-4.27v-.05l-.7-.64A1.11 1.11 0 0 1 6.33 7c0-.58.28-1.3.84-2.18C7.93 3.6 8.7 3 9.46 3c1.03 0 1.89.49 2.56 1.45c.38.56-.03.65-1.24.26c-.98-.38-1.78-.06-2.4.96l.02.1l1.31 1h.06c1.64-.57 2.82-.86 3.55-.84a5.5 5.5 0 0 0-.28.86a32.4 32.4 0 0 1-.36 1.14l-.14.44l-.45.05c-2.04.28-3.5.84-4.37 1.67a2.5 2.5 0 0 0-.83 1.78"/>`),
		g.Group(children),
	)
}