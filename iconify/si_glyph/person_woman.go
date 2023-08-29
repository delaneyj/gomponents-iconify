package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonWoman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.677 11c-2.843 2.147-.724 3.424-3.653 3.424c-2.932 0-.604-1.191-3.76-3.398c-4.908 1.381-4.226 4.884-4.226 4.884h15.875c-.001 0 .651-3.921-4.236-4.91zm-.103-4.002l-.575 1.88L13.883 9s-1.153-.85-1.309-2.002zm-7.18.025l.538 1.893L4.045 9s1.169-.826 1.349-1.977zm6.535-1.977c0 1.68-1.323 4.895-2.954 4.895c-1.632 0-2.954-3.215-2.954-4.895C6.021 3.363 7.343 2 8.975 2c1.63.001 2.954 1.364 2.954 3.046z"/>`),
		g.Group(children),
	)
}