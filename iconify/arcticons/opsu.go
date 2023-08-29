package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opsu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.457 3.992l7.83 8.923a1.4 1.4 0 0 0 1.217.467l11.79-1.392a1.4 1.4 0 0 1 1.367 2.106L37.595 24.3a1.4 1.4 0 0 0-.069 1.302l4.967 10.783a1.4 1.4 0 0 1-1.58 1.95l-11.58-2.616a1.4 1.4 0 0 0-1.258.338l-8.72 8.056a1.4 1.4 0 0 1-2.344-.9L15.92 31.39a1.4 1.4 0 0 0-.71-1.093L4.854 24.494a1.4 1.4 0 0 1 .132-2.507l10.906-4.69a1.4 1.4 0 0 0 .82-1.013l2.32-11.643a1.4 1.4 0 0 1 2.425-.65Z"/><rect width="6.355" height="8.42" x="20.689" y="21.297" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.177" transform="rotate(-15 23.866 25.507)"/><circle cx="30.572" cy="27.84" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.794 17.472l2.066 7.711"/>`),
		g.Group(children),
	)
}