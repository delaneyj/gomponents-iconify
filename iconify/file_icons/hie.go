package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 386.728L87.153 256L0 125.272h65.364L152.517 256L65.364 386.728H0zm87.153 0L174.306 256L87.153 125.272h65.364L326.82 386.728h-65.364l-54.468-81.704l-54.471 81.704H87.153zm395.795-130.74L512 212.414l-145.254-.004v43.579h116.202zm-58.101 65.364l29.049-43.575l-87.15-.003v43.578h58.101zm29.049 65.364l29.052-43.575l-116.202-.003v43.578h87.15zm-148.864-65.352l43.575 65.364l.003-174.318h-43.578v108.954zm43.578-127.092l.001-43.575l-43.579-.003v43.578h43.578z"/>`),
		g.Group(children),
	)
}