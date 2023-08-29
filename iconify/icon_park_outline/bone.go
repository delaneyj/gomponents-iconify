package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M43.523 17.315c0-4.243-3.536-4.243-6.01-6.718c-2.475-2.475-2.475-6.01-6.718-6.01c-2.829 0-5.303 3.182-5.303 5.303s.707 3.536 2.121 4.95L14.885 27.568c-1.414-1.414-2.828-2.122-4.95-2.122c-2.121 0-5.303 2.475-5.303 5.304c-.354 4.596 3.89 4.596 6.01 6.717c2.122 2.122 2.122 6.364 6.718 6.01c2.828 0 5.303-3.181 5.303-5.303c0-2.121-.707-3.535-2.121-4.95L33.27 20.498c1.414 1.414 2.828 2.121 4.95 2.121c2.12 0 5.303-2.475 5.303-5.303Z"/>`),
		g.Group(children),
	)
}