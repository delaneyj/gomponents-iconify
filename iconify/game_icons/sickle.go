package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sickle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M257.563 25.156c353.16 276.87 16.918 408.895-87.875 293.25l-40.75 37.125l50.812 50.345c217.562 181.363 524.73-252.058 77.813-380.72zM110.75 364.28c-5.525 1.065-8.975 2.957-11.313 5.25c-1.956 1.922-3.248 4.556-4.25 7.564l55.188 52.844c5.468-1.008 9.264-2.796 11.28-4.688c1.997-1.872 3.095-3.864 3.095-7.53l-54-53.44zm-24.72 30.314L30.407 445.28c-16.737 27 14.693 61.198 51.093 44.66l51.53-50.282l-47-45.062z"/>`),
		g.Group(children),
	)
}